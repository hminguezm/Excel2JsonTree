package main

import (
	"encoding/json"
	"log"
	"time"
	"github.com/onkarvhanumante/Excel2JsonTree"
	"fmt"
	"errors"
)

type MenuInfo struct {
	ParentName  		string     `json:"parentName"`
	DisplayName 		string     `json:"displayName"`
	MenuDescription     string     `json:"menuDescription"`
	Ordinal     		string     `json:"ordinal"`
	ValidTill   		time.Time  `json:"validTill"`
	RoleList    		[]string   `json:"roleList"`
	HasChild    		bool       `json:"hasChild"`
	MenuList   			[]MenuInfo `json:"menuList"`
	OptionLevel       	int        `json:"optionLevel"`
	CreatedBy			string	   `json:"createdBy"`
}

func main() {

	// read excel
	excelFile, err := xlsToJson.ReadExcel("./Menu.xlsx")
	if err != nil {
		log.Fatal("error in reading excel : ", err)
	}

	// check sheet
	if len(excelFile.Sheets) == 0 {
		log.Fatal("main : excel has no sheet")
	}

	// s1 : get variable holder pointer
	variableHolderObjPtr := xlsToJson.GetVariableHolderPtr()

	// s2 : set name of column(s) to be refered for generating json
	err = variableHolderObjPtr.SetNodesColumnNameMap([]string{"level"})
	if err != nil {
		log.Fatal("main : error in setting SetNodesColumnNameMap : ", err)
	}

	// s3 : set AttributeMappingMap
	// AttributeMappingMap : key - excel header and value - struct keys
	// set inputs and functions to be performed during json generation  (optional)
	
	structXslMap := make(map[string]string)
	structXslMap["level"] = "DisplayName"
	structXslMap["description"] = "MenuDescription"
	structXslMap["ordinal"] = "Ordinal"
	structXslMap["validTill"] = "ValidTill"
	structXslMap["roleList"] = "RoleList"
	structXslMap["createrID"] = "CreatedBy"

	// SetInputParametersToUserDefinedFunc - takes two arguements
	// arg1 : key name to access
	// arg2 : value
	
	IdDatabase := make(map[string]string)
	IdDatabase["1"] = "user1"
	IdDatabase["2"] = "user2"
	variableHolderObjPtr.SetInputParametersToUserDefinedFunc("IdDatabase", IdDatabase)
	
	// SetKeyFunctionMap - takes two arguements
	// arg1 : struct key name
	// arg2 : function to be performed	
	variableHolderObjPtr.SetKeyFunctionMap("ValidTill", generateValidTillAttributeValue)
	variableHolderObjPtr.SetKeyFunctionMap("ParentName", assignParentName)
	variableHolderObjPtr.SetKeyFunctionMap("HasChild", assignCurrentNodeChildStatus)
	variableHolderObjPtr.SetKeyFunctionMap("CreatedBy", assignCreatedBy)
	variableHolderObjPtr.SetKeyFunctionMap("OptionLevel", assignOptionLevel)
	err = variableHolderObjPtr.SetAttributeMappingMap(structXslMap)
	if err != nil {
		log.Fatal("main : error in setting SetAttributeMappingMap : ", err)
	}

	// s4: generate json
	// takes three arguements
	// arg1 : excel sheet
	// arg2 : struct
	// arg3 : struct child node key name
	jsonTree, err := variableHolderObjPtr.ConvertXslToJSON(excelFile.Sheets[0], MenuInfo{}, "MenuList")
	if err != nil {
		log.Fatal("ConvertXslToJSON : error in converting excel to json : ", err)
	}
	
	a := jsonTree.(MenuInfo)
	b, _ := json.Marshal(a.MenuList)
	fmt.Println("Json:")
	fmt.Println(string(b))
	
}

// sample functions
// each function receives 3 args
// arg1 : input set using SetInputParametersToUserDefinedFunc function
// arg2 : excel cell value on which function is supposed to perform operation
// arg3 : variable holder instance to call built in public methods - GetCurrentNodeLevel(), GetParentNodeForChild(), GetCurrentNode(), GetCurrentNodeChildStatus()

var generateValidTillAttributeValue xlsToJson.UserDefinedFunction = func(inputMap map[string]interface{}, excelCellValue string, v *xlsToJson.VariableHolder) (interface{}, error) {
	var output interface{}
	parsedDate, err := time.Parse("02/01/2006", excelCellValue)
	if err != nil {
		return output, err
	}
	output = parsedDate
	return output, nil
}

var assignOptionLevel xlsToJson.UserDefinedFunction = func(inputMap map[string]interface{}, excelCellValue string, v *xlsToJson.VariableHolder) (interface{}, error) {
	var output interface{}
	output = v.GetCurrentNodeLevel()
	return output, nil
}

var assignParentName xlsToJson.UserDefinedFunction = func(inputMap map[string]interface{}, excelCellValue string, v *xlsToJson.VariableHolder) (interface{}, error) {
	var output interface{}
	parent := v.GetParentNodeForChild()
	if parent != nil {
		parent1 := parent.(MenuInfo)
		output = parent1.DisplayName
	} else {
		output = "NA"
	}

	return output, nil
}

var assignCurrentNodeChildStatus xlsToJson.UserDefinedFunction = func(inputMap map[string]interface{}, excelCellValue string, v *xlsToJson.VariableHolder) (interface{}, error) {
	var output interface{}
	output = v.GetCurrentNodeChildrenStatus()
	return output, nil
}

var assignCreatedBy xlsToJson.UserDefinedFunction = func(inputMap map[string]interface{}, excelCellValue string, v *xlsToJson.VariableHolder) (interface{}, error) {
	var output interface{}
	idDb := inputMap["IdDatabase"].(map[string]string)
	val, ok := idDb[excelCellValue]
	if ok {
		output = val
	} else {
		return nil, errors.New("Id not present in db")
	}
	return output, nil
}
