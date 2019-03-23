package xlsToJson

import (
	"reflect"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

// appendChildNode - appends child node in parent
func (v *VariableHolder) appendChildNode(parent, child reflect.Value) reflect.Value {
	if parent.Kind() == reflect.Ptr {
		parent = parent.Elem()
	}
	if child.Kind() == reflect.Ptr {
		child = child.Elem()
	}
	parentChildrenList := parent.FieldByName(v.childrenArrayKey)
	parentChildrenList.Set(reflect.Append(parentChildrenList, child))
	return parent
}

// deReferenceNode -  derefernces if reflect value is of type pointer
// to be used during appending
func deReferenceNode(node reflect.Value) reflect.Value {
	if node.Kind() == reflect.Ptr {
		return node.Elem()
	}
	return node
}

// createNewNode - create new node based on value received
// create child using reflect
// iterate over cells in current row and get type and struct key for each cell value and assign it to created child node
// iterate overkeyFunctionMapKey and get type and struct key for each jey value and assign it to created child node
func (v *VariableHolder) createNewNode(currentRow *xlsx.Row, cellIndex int) (reflect.Value, error) {
	childNode := reflect.New(v.childrenArrayEntryType)
	funcExecutedKeyMap := make(map[string]bool)
	for i, cell := range currentRow.Cells {
		if len(cell.String()) != 0 {
			valueType, structKey := v.getChildNodeValueTypeAndStructKeyForAssigningExcelData(i)
			var valReturnByUserDefinedFunc interface{}
			var err error
			val, ok := v.keyFunctionMap[structKey]
			if ok {
				valReturnByUserDefinedFunc, err = val(v.methodsInputParameterMap, cell.String(), v)
				if err != nil {
					return childNode, err
				}
				if reflect.ValueOf(valReturnByUserDefinedFunc).Type().String() != valueType.String() {
					return childNode, generateError(ERR_CODE_USERDEFINED_FUNCTION_RETURN_VALUE_AND_STRUCT_KEY_TYPE_MISMATCH + " for struct key : " + structKey)
				}
			}
			switch valueType.String() {
			case STRING:
				if ok {
					childNode.Elem().FieldByName(structKey).SetString(valReturnByUserDefinedFunc.(string))
					funcExecutedKeyMap[structKey] = true
				} else {
					childNode.Elem().FieldByName(structKey).SetString(cell.String())
				}
			case INT, INT8, INT32, INT64:
				if ok {
					if valueType.String() == INT {
						childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int)))
					} else if valueType.String() == INT8 {
						childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int8)))
					} else if valueType.String() == INT32 {
						childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int32)))
					} else if valueType.String() == INT64 {
						childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int64)))
					}
					funcExecutedKeyMap[structKey] = true
				} else {
					intVal, _ := cell.Int()
					childNode.Elem().FieldByName(structKey).SetInt(int64(intVal))
				}
			case BOOLEAN:
				if ok {
					childNode.Elem().FieldByName(structKey).SetBool(valReturnByUserDefinedFunc.(bool))
					funcExecutedKeyMap[structKey] = true
				} else {
					childNode.Elem().FieldByName(structKey).SetBool(cell.Bool())
				}
			case INT8_SLICE, INT_SLICE, STRING_SLICE, BOOLEAN_SLICE, INT32_SLICE, INT64_SLICE, FLOAT32_SLICE, FLOAT64_SLICE:
				if ok {
					var slice reflect.Value
					if valueType.String() == INT8_SLICE {
						v := valReturnByUserDefinedFunc.([]int8)
						slice = reflect.MakeSlice(reflect.TypeOf([]int8{}), len(v), len(v))
					} else if valueType.String() == INT_SLICE {
						v := valReturnByUserDefinedFunc.([]int)
						slice = reflect.MakeSlice(reflect.TypeOf([]int{}), len(v), len(v))
					} else if valueType.String() == STRING_SLICE {
						v := valReturnByUserDefinedFunc.([]string)
						slice = reflect.MakeSlice(reflect.TypeOf([]string{}), len(v), len(v))
					} else if valueType.String() == BOOLEAN_SLICE {
						v := valReturnByUserDefinedFunc.([]bool)
						slice = reflect.MakeSlice(reflect.TypeOf([]bool{}), len(v), len(v))
					} else if valueType.String() == INT32_SLICE {
						v := valReturnByUserDefinedFunc.([]int32)
						slice = reflect.MakeSlice(reflect.TypeOf([]int32{}), len(v), len(v))
					} else if valueType.String() == INT64_SLICE {
						v := valReturnByUserDefinedFunc.([]int64)
						slice = reflect.MakeSlice(reflect.TypeOf([]int64{}), len(v), len(v))
					} else if valueType.String() == FLOAT32_SLICE {
						v := valReturnByUserDefinedFunc.([]float32)
						slice = reflect.MakeSlice(reflect.TypeOf([]float32{}), len(v), len(v))
					} else if valueType.String() == FLOAT64_SLICE {
						v := valReturnByUserDefinedFunc.([]float64)
						slice = reflect.MakeSlice(reflect.TypeOf([]float64{}), len(v), len(v))
					}
					childNode.Elem().FieldByName(structKey).Set(slice)
					funcExecutedKeyMap[structKey] = true
				} else {
					childNode.Elem().FieldByName(structKey).Set(getSliceOfTypeReflect(v.getStringList(cell.String()), valueType.String()))
				}
			case FLOAT32, FLOAT64:
				if ok {
					if valueType.String() == FLOAT32 {
						childNode.Elem().FieldByName(structKey).SetFloat(float64(valReturnByUserDefinedFunc.(float32)))
					} else if valueType.String() == FLOAT64 {
						childNode.Elem().FieldByName(structKey).SetFloat(valReturnByUserDefinedFunc.(float64))
					}
					funcExecutedKeyMap[structKey] = true
				} else {
					floatValue, _ := strconv.ParseFloat(cell.String(), 64)
					childNode.Elem().FieldByName(structKey).SetFloat(floatValue)
				}
			case TIME_DOT_TIME:
				if ok {
					childNode.Elem().FieldByName(structKey).Set(reflect.ValueOf(valReturnByUserDefinedFunc.(time.Time)))
					funcExecutedKeyMap[structKey] = true
				} else {
					parsedDate, err := time.Parse(v.dateFormatLayout, cell.String())
					if err != nil {
						return reflect.New(v.childrenArrayEntryType), generateError(ERR_CODE_DATE_FORMAT_PARSING_ERROR)
					}
					childNode.Elem().FieldByName(structKey).Set(reflect.ValueOf(parsedDate))
				}
			default:
				if valueType.Kind().String() != STRUCT && valueType.Kind().String() != SLICE {
					if ok {
						childNode.Elem().FieldByName(structKey).SetString(valReturnByUserDefinedFunc.(string))
						funcExecutedKeyMap[structKey] = true
					} else {
						childNode.Elem().FieldByName(structKey).SetString(cell.String())
					}
				}
			}
		}
	}
	KeyFunctionMapKeyList := reflect.ValueOf(v.keyFunctionMap).MapKeys()
	for i := 0; i < len(KeyFunctionMapKeyList); i++ {
		if funcExecutedKeyMap[KeyFunctionMapKeyList[i].String()] {
			continue
		}
		valueType, structKey, err := v.getChildNodeValueTypeAndStructKeyForAssigningDataFromUserDefinedFunction(KeyFunctionMapKeyList[i].String())
		if err != nil {
			return childNode, err
		}
		valReturnByUserDefinedFunc, err := v.keyFunctionMap[KeyFunctionMapKeyList[i].String()](v.methodsInputParameterMap, "", v)
		if err != nil {
			return childNode, err
		}
		if reflect.ValueOf(valReturnByUserDefinedFunc).Type().String() != valueType.String() {
			return childNode, generateError(ERR_CODE_USERDEFINED_FUNCTION_RETURN_VALUE_AND_STRUCT_KEY_TYPE_MISMATCH + " for struct key : " + structKey)
		}
		caseSatisfied := false
		switch valueType.String() {
		case STRING:
			childNode.Elem().FieldByName(structKey).SetString(valReturnByUserDefinedFunc.(string))
			caseSatisfied = true
		case INT, INT8, INT32, INT64:
			if valueType.String() == INT {
				childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int)))
			} else if valueType.String() == INT8 {
				childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int8)))
			} else if valueType.String() == INT32 {
				childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int32)))
			} else if valueType.String() == INT64 {
				childNode.Elem().FieldByName(structKey).SetInt(int64(valReturnByUserDefinedFunc.(int64)))
			}
			caseSatisfied = true
		case BOOLEAN:
			childNode.Elem().FieldByName(structKey).SetBool(valReturnByUserDefinedFunc.(bool))
			caseSatisfied = true
		case INT8_SLICE, INT_SLICE, STRING_SLICE, BOOLEAN_SLICE, INT32_SLICE, INT64_SLICE, FLOAT32_SLICE, FLOAT64_SLICE:
			var slice reflect.Value
			if valueType.String() == INT8_SLICE {
				v := valReturnByUserDefinedFunc.([]int8)
				slice = reflect.MakeSlice(reflect.TypeOf([]int8{}), len(v), len(v))
			} else if valueType.String() == INT_SLICE {
				v := valReturnByUserDefinedFunc.([]int)
				slice = reflect.MakeSlice(reflect.TypeOf([]int{}), len(v), len(v))
			} else if valueType.String() == STRING_SLICE {
				v := valReturnByUserDefinedFunc.([]string)
				slice = reflect.MakeSlice(reflect.TypeOf([]string{}), len(v), len(v))
			} else if valueType.String() == BOOLEAN_SLICE {
				v := valReturnByUserDefinedFunc.([]bool)
				slice = reflect.MakeSlice(reflect.TypeOf([]bool{}), len(v), len(v))
			} else if valueType.String() == INT32_SLICE {
				v := valReturnByUserDefinedFunc.([]int32)
				slice = reflect.MakeSlice(reflect.TypeOf([]int32{}), len(v), len(v))
			} else if valueType.String() == INT64_SLICE {
				v := valReturnByUserDefinedFunc.([]int64)
				slice = reflect.MakeSlice(reflect.TypeOf([]int64{}), len(v), len(v))
			} else if valueType.String() == FLOAT32_SLICE {
				v := valReturnByUserDefinedFunc.([]float32)
				slice = reflect.MakeSlice(reflect.TypeOf([]float32{}), len(v), len(v))
			} else if valueType.String() == FLOAT64_SLICE {
				v := valReturnByUserDefinedFunc.([]float64)
				slice = reflect.MakeSlice(reflect.TypeOf([]float64{}), len(v), len(v))
			}
			childNode.Elem().FieldByName(structKey).Set(slice)
			caseSatisfied = true
		case FLOAT32, FLOAT64:
			if valueType.String() == FLOAT32 {
				childNode.Elem().FieldByName(structKey).SetFloat(float64(valReturnByUserDefinedFunc.(float32)))
			} else if valueType.String() == FLOAT64 {
				childNode.Elem().FieldByName(structKey).SetFloat(valReturnByUserDefinedFunc.(float64))
			}
			caseSatisfied = true
		case TIME_DOT_TIME:
			childNode.Elem().FieldByName(structKey).Set(reflect.ValueOf(valReturnByUserDefinedFunc.(time.Time)))
			caseSatisfied = true
		default:
			if valueType.Kind().String() != STRUCT && valueType.Kind().String() != SLICE {
				childNode.Elem().FieldByName(structKey).SetString(valReturnByUserDefinedFunc.(string))
				caseSatisfied = true
			}
		}
		if !caseSatisfied {
			switch valueType.Kind().String() {
			case STRUCT, SLICE:
				v := reflect.ValueOf(valReturnByUserDefinedFunc)
				v = deReferenceNode(v)
				childNode.Elem().FieldByName(structKey).Set(v)
			}
		}
	}
	return childNode, nil
}

// GenerateTree - create tree
func (v *VariableHolder) generateTree() (reflect.Value, error) {
	nodeListV1 := reflect.New(v.childrenArrayEntryType)
	current := -1
	previous := -1
	var err error
	for i := v.excelHeaderRowIndexNo + 1; i < len(v.currentXls.Rows); i++ {
		v.setCurrentRowIndexNoBeingProcessed(i)
		for j := 0; j < len(v.currentXls.Rows[i].Cells); j++ {
			if v.currentXls.Rows[i].Cells[j].String() != "" && v.nodeColumnNameMap[v.excelHeaderRow.Cells[j].String()] {
				current = j
				v.setCurrentNodeLevel(current)
				if current == 0 {
					v.setCurrentNodeLevelFlag(APPEND_NODE_AT_START_LEVEL)
					nodeChild, err := v.createNewNode(v.currentXls.Rows[i], current)
					if err != nil {
						return reflect.Value{}, err
					}
					nodeChild = deReferenceNode(nodeChild)
					v.updateCurrentNode(nodeChild.Interface())
					v.updatelevelWiseNodeMapEntry(current, nodeChild.Interface())
					nodeListV1 = v.appendChildNode(nodeListV1, nodeChild)
				} else if current > previous {
					v.setCurrentNodeLevelFlag(APPEND_NODE_AT_NEXT_LEVEL)
					nodeListV1, err = v.addNodeToNextLevelV1(current, nodeListV1, v.currentXls.Rows[i])
					if err != nil {
						return reflect.Value{}, err
					}
				} else if current < previous {
					v.setCurrentNodeLevelFlag(APPEND_NODE_AT_PREVIOUS_LEVEL)
					nodeListV1, err = v.addNodeToPreviousLevelV1(current, nodeListV1, v.currentXls.Rows[i])
					if err != nil {
						return reflect.Value{}, err
					}
				} else if current == previous {
					v.setCurrentNodeLevelFlag(APPEND_NODE_AT_SAME_LEVEL)
					nodeListV1, err = v.addNodeToSameLevelV1(current, nodeListV1, v.currentXls.Rows[i])
					if err != nil {
						return reflect.Value{}, err
					}
				}
			}
		}
		previous = current
	}
	return nodeListV1, nil
}

// ConvertXslToJSON - converts the data from sheet to JSON
func (v *VariableHolder) ConvertXslToJSON(sheet *xlsx.Sheet, structReference interface{}, childrenArrayKey string) (interface{}, error) {
	if reflect.TypeOf(structReference).Kind() != reflect.Struct {
		return nil, generateError(ERR_CODE_INTERFACE_RECEIVED_IS_NOT_STRUCT)
	}
	if len(childrenArrayKey) == 0 {
		return nil, generateError(ERR_CODE_CHILDREN_ARRAY_KEY_NOT_SET)
	}
	err := v.checkAndSetCurrentXlsInfo(sheet, childrenArrayKey)
	if err != nil {
		return nil, err
	}
	v.setKeyToAppendChildrenNodes(childrenArrayKey)
	err = v.setStructInfoToGenerateTree(structReference)
	if err != nil {
		return nil, generateError(ERR_CODE_UNABLE_TO_SET_STRUCT_INFO)
	}
	err = v.checkHeaderRowColumnWithMaps()
	if err != nil {
		return nil, err
	}
	jsonTree, err := v.generateTree()
	if err != nil {
		return nil, err
	}
	return jsonTree.Interface(), nil
}
