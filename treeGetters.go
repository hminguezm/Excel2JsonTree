package xlsToJson

import (
	"reflect"
	"strconv"
	"strings"
)

// getChildNodeValueTypeAndStructKeyForAssigningExcelData - returns the value type for the corresponding cell index and struct key
func (v *VariableHolder) getChildNodeValueTypeAndStructKeyForAssigningExcelData(cellIndex int) (reflect.Type, string) {
	return v.structKeyValueMap[v.attributeMappingMap[v.excelHeaderRow.Cells[cellIndex].String()]], v.attributeMappingMap[v.excelHeaderRow.Cells[cellIndex].String()]
}

// getChildNodeValueTypeAndStructKeyForAssigningDataFromUserDefinedFunction - returns the value type for the corresponding cell index and struct key
func (v *VariableHolder) getChildNodeValueTypeAndStructKeyForAssigningDataFromUserDefinedFunction(key string) (reflect.Type, string, error) {
	value, isPresent := v.structKeyValueMap[key]
	if isPresent {
		return value, key, nil
	}
	return value, "", generateError(ERR_CODE_KEY_NOT_PRESENT_IN_STRUCT + ":" + key)
}

// getLastNodeFromChildrenListOfNode - returns the last child from node children list
func (v *VariableHolder) getLastNodeFromChildrenListOfNode(node reflect.Value) reflect.Value {
	parentChildrenList := node.FieldByName(v.childrenArrayKey)
	if parentChildrenList.Len() == 0 {
		return reflect.New(v.childrenArrayEntryType)
	}
	return parentChildrenList.Index(parentChildrenList.Len() - 1)
}

// getSliceOfTypeReflect - returns the slice of received reflect type
func getSliceOfTypeReflect(stringList []string, reflectType string) reflect.Value {
	switch reflectType {
	case INT_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]int{}), len(stringList), len(stringList))
		for i, str := range stringList {
			intValue, _ := strconv.Atoi(str)
			slice.Index(i).Set(reflect.ValueOf(int8(intValue)))
		}
		return slice
	case INT8_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]int{}), len(stringList), len(stringList))
		for i, str := range stringList {
			intValue, _ := strconv.Atoi(str)
			slice.Index(i).Set(reflect.ValueOf(int(intValue)))
		}
		return slice
	case INT32_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]int32{}), len(stringList), len(stringList))
		for i, str := range stringList {
			intValue, _ := strconv.Atoi(str)
			slice.Index(i).Set(reflect.ValueOf(int32(intValue)))
		}
		return slice
	case INT64_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]int64{}), len(stringList), len(stringList))
		for i, str := range stringList {
			intValue, _ := strconv.Atoi(str)
			slice.Index(i).Set(reflect.ValueOf(int64(intValue)))
		}
		return slice
	case FLOAT32_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]float32{}), len(stringList), len(stringList))
		for i, str := range stringList {
			floatValue, _ := strconv.ParseFloat(str, 32)
			slice.Index(i).Set(reflect.ValueOf(floatValue))
		}
		return slice
	case FLOAT64_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]float64{}), len(stringList), len(stringList))
		for i, str := range stringList {
			floatValue, _ := strconv.ParseFloat(str, 64)
			slice.Index(i).Set(reflect.ValueOf(floatValue))
		}
		return slice
	case STRING_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]string{}), len(stringList), len(stringList))
		for i, str := range stringList {
			slice.Index(i).Set(reflect.ValueOf(str))
		}
		return slice
	case BOOLEAN_SLICE:
		slice := reflect.MakeSlice(reflect.TypeOf([]bool{}), len(stringList), len(stringList))
		for i, str := range stringList {
			b, _ := strconv.ParseBool(str)
			slice.Index(i).Set(reflect.ValueOf(b))
		}
		return slice
	default:
		// if unknown type then make slice of string
		slice := reflect.MakeSlice(reflect.TypeOf([]string{}), len(stringList), len(stringList))
		for i, str := range stringList {
			slice.Index(i).Set(reflect.ValueOf(str))
		}
		return slice
	}
}

// getStringList - prepares the string list
func (v *VariableHolder) getStringList(str string) []string {
	str = strings.NewReplacer("[", EMPTY_STR, "]", EMPTY_STR, "\"", EMPTY_STR).Replace(str)
	return strings.Split(str, v.listStrSeperator)
}
