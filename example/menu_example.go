package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/onkarvhanumante/Excel2JsonTree"
)

type MenuInfo struct {
	ParentName  string     `json:"parentName"`
	DisplayName string     `json:"displayName" excel:"level"`
	MenuType    string     `json:"menuType" excel:"menuType"`
	Menu        string     `json:"menu" excel:"menu"`
	Ordinal     string     `json:"ordinal" excel:"ordinal"`
	ValidTill   time.Time  `json:"validTill" excel:"ValidTill"`
	RoleList    []string   `json:"roleList" excel:"RoleList"`
	HasChild    bool       `json:"hasChild"`
	MenuList    []MenuInfo `json:"menuList"`
	Level       int        `json:"level"`
}

func main() {

	// read excel
	excelFile, err := xlsToJson.ReadExcel("C:/Users/onkarh/Desktop/xls/Menu.xlsx")
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
	// (optional)
	// SetInputParametersToUserDefinedFunc - takes two arguements
	// arg1 : struct key name
	// arg2 : function to be performed

	variableHolderObjPtr.SetKeyFunctionMap("ValidTill", generateValidTillAttributeValue)
	variableHolderObjPtr.SetKeyFunctionMap("Level", generateLevel)
	variableHolderObjPtr.SetKeyFunctionMap("ParentName", generateParentName)
	variableHolderObjPtr.SetKeyFunctionMap("HasChild", assignCurrentNodeChildStatus)

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
	parsedDate, err := time.Parse("01-02-06", excelCellValue)
	if err != nil {
		logginghelper.LogError("generateValidTillAttributeValue : error in parsing date : ", err)
		return output, err
	}

	output = parsedDate
	return output, nil
}

var generateLevel xlsToJson.UserDefinedFunction = func(inputMap map[string]interface{}, excelCellValue string, v *xlsToJson.VariableHolder) (interface{}, error) {
	var output interface{}
	output = v.GetCurrentNodeLevel()
	return output, nil
}

var generateParentName xlsToJson.UserDefinedFunction = func(inputMap map[string]interface{}, excelCellValue string, v *xlsToJson.VariableHolder) (interface{}, error) {
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
