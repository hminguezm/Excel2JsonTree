package xlsToJson

import (
	"reflect"
	"strings"
)

// updatelevelWiseNodeMapEntry - adds/ updates entry in levelWiseNodeMap
// always stores latest entry of node at a given level
func (v *VariableHolder) updatelevelWiseNodeMapEntry(level int, node interface{}) {
	v.levelWiseNodeMap[level] = node
}

// updateCurrentNode - adds/ updates entry in currentNode
// always stores latest entry of node at given level
func (v *VariableHolder) updateCurrentNode(node interface{}) {
	v.currentNode = node
}

// setCurrentNodeLevel - sets current node level
func (v *VariableHolder) setCurrentNodeLevel(level int) {
	v.currentNodeLevel = level
}

// setCurrentNodeLevelFlag - sets the current node level
func (v *VariableHolder) setCurrentNodeLevelFlag(level string) {
	v.currentNodeLevelFlag = level
}

// setStructInfoToGenerateTree - checks struct used to generate json
// checks keys of struct received and store there type
func (v *VariableHolder) setStructInfoToGenerateTree(referenceStruct interface{}) error {

	// get number of keys received in reference struct
	t := reflect.TypeOf(referenceStruct)
	val := reflect.Indirect(reflect.ValueOf(referenceStruct))

	// maintain map of referenceStruct key as map key and referenceStruct key's value as map value
	for i := 0; i < t.NumField(); i++ {
		v.structKeyValueMap[val.Type().Field(i).Name] = val.Type().Field(i).Type
	}

	// set type of elements to be appended in children slice
	v.setChildrenArrayEntryType(v.structKeyValueMap[v.childrenArrayKey].Elem())
	return nil
}

// setChildrenArrayEntryType - used to set childrenArrayEntryType
func (v *VariableHolder) setChildrenArrayEntryType(entryType reflect.Type) {
	v.childrenArrayEntryType = entryType
	// childrenArrayEntryType = entryType
}

// SetAttributeMappingMap - sets the data in attributeMappingMap
// structXslMap - key :  excel column name, value - struct key name
// sends error if map received is empty
// sends error if key received in maps are also present in nodeColumnNameMap
func (v *VariableHolder) SetAttributeMappingMap(structXslMap map[string]string) error {

	// check if map received is empty
	structXslMapKeys := reflect.ValueOf(structXslMap).MapKeys()
	if len(structXslMapKeys) == 0 {
		return generateError(ERR_CODE_MAP_RECEIVED_EMPTY)
	}

	// check if nodeColumnNameMap keys were present in structXslMap
	nodeColumnKeyPresentCount := 0
	nodeColumnNameMapKeys := reflect.ValueOf(v.nodeColumnNameMap).MapKeys()
	for _, key := range nodeColumnNameMapKeys {
		if len(structXslMap[key.String()]) != 0 {
			nodeColumnKeyPresentCount++
		}
	}
	if nodeColumnKeyPresentCount != len(v.nodeColumnNameMap) {
		return generateError(ERR_CODE_NODE_COLUMN_NAME_NOT_PRESENT_IN_MAP)
	}
	// assign map
	v.attributeMappingMap = structXslMap
	return nil
}

// SetInputParametersToUserDefinedFunc - sets the input  in  methodsInputParameterMap that can be used in user defined function
func (v *VariableHolder) SetInputParametersToUserDefinedFunc(key string, value interface{}) {
	v.methodsInputParameterMap[key] = value
}

// SetKeyFunctionMap - sets the function to be executed on particular key
func (v *VariableHolder) SetKeyFunctionMap(structKey string, bindingFunction UserDefinedFunction) {
	v.keyFunctionMap[structKey] = bindingFunction
}

// SetNodesColumnNameMap - sets the name of column to be searched
// sends error if map received is empty
func (v *VariableHolder) SetNodesColumnNameMap(nodeNameColumnList []string) error {
	if len(nodeNameColumnList) == 0 {
		return generateError(ERR_CODE_EMPTY_LIST_RECEIVED)
	}
	for _, nodeName := range nodeNameColumnList {
		v.nodeColumnNameMap[strings.ToLower(nodeName)] = true
	}
	return nil
}

// setKeyToAppendChildrenNodes - assgin value received to childrenArrayKey
// used as key to append node as child
func (v *VariableHolder) setKeyToAppendChildrenNodes(keyName string) {
	v.childrenArrayKey = keyName
}

// setCurrentRowIndexNoBeingProcessed - index number of row being processed
func (v *VariableHolder) setCurrentRowIndexNoBeingProcessed(indexNo int) {
	v.currentRowNo = indexNo
}

// SetListStringSeperator - used to assgin the seperator to split the string
func (v *VariableHolder) SetListStringSeperator(seperator string) {
	v.listStrSeperator = seperator
}

// SetExcelHeaderRowIndexNo - index no of excel header row
func (v *VariableHolder) SetExcelHeaderRowIndexNo(indexNo int) {
	v.excelHeaderRowIndexNo = indexNo
}
