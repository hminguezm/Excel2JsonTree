package xlsToJson

import "reflect"

// GetParentNodeForChild - returns the parent for current node
func (v *VariableHolder) GetParentNodeForChild() interface{} {
	if v.currentNodeLevelFlag == APPEND_NODE_AT_START_LEVEL {
		return nil
	}
	return v.levelWiseNodeMap[v.GetCurrentNodeLevel()-1]
}

// GetCurrentNodeLevel - returns the current node level
func (v *VariableHolder) GetCurrentNodeLevel() int {
	return v.currentNodeLevel
}

// GetCurrentNode - returns the current node
func (v *VariableHolder) GetCurrentNode() interface{} {
	return v.currentNode
}

// GetCurrentNodeChildrenStatus - return true if current node has child
func (v *VariableHolder) GetCurrentNodeChildrenStatus() bool {
	// current row is not last row
	if v.currentRowNo+1 != v.currentXlsRowCount {
		if len(v.currentXls.Rows[v.currentRowNo+1].Cells[v.currentNodeLevel+1].String()) != 0 && v.nodeColumnNameMap[v.excelHeaderRow.Cells[v.currentNodeLevel+1].String()] {
			return true
		}
	}
	return false
}

// GetVariableHolderPtr - returns the VariableHolder instance
func GetVariableHolderPtr() *VariableHolder {
	v := VariableHolder{}
	v.levelWiseNodeMap = make(map[int]interface{})
	v.nodeColumnNameMap = make(map[string]bool)
	v.attributeMappingMap = make(map[string]string)
	v.structKeyValueMap = make(map[string]reflect.Type)
	v.keyFunctionMap = make(map[string]UserDefinedFunction)
	v.methodsInputParameterMap = make(map[string]interface{})
	v.dateFormatLayout = DEFAULT_DATE_FORMAT_LAYOUT
	v.listStrSeperator = DEFAULT_LIST_STRING_SEPERATOR
	v.excelHeaderRowIndexNo = DEFAULT_HEADER_ROW_INDEX
	return &v
}
