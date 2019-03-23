package xlsToJson

import (
	"reflect"
)

// checkHeaderRowColumnWithMaps - check header value with Maps
func (v *VariableHolder) checkHeaderRowColumnWithMaps() error {

	// maintain excel header row unique key
	excelHeaderRowUniqueKeyMap := make(map[string]bool)

	// for _, cell := range excelHeaderRow.Cells {
	for _, cell := range v.excelHeaderRow.Cells {
		if !excelHeaderRowUniqueKeyMap[cell.String()] && len(cell.String()) != 0 {
			excelHeaderRowUniqueKeyMap[cell.String()] = true
		}
	}

	// attributeMappingMap - check if number of parameter set in attribute map are correct
	if len(v.attributeMappingMap) != len(excelHeaderRowUniqueKeyMap) {
		return generateError(ERR_CODE_KEYS_COUNT_MISMATCHED_IN_INPUT_SEND_TO_SETATTRIBUTEMAPPINGMAP)
	}

	// maintain attribut map key list
	attributeMappingMapkeyList := reflect.ValueOf(v.attributeMappingMap).MapKeys()
	for i := 0; i < len(attributeMappingMapkeyList); i++ {
		if !excelHeaderRowUniqueKeyMap[attributeMappingMapkeyList[i].String()] {
			return generateError(ERR_CODE_KEYS_MISMATCHED_IN_ATTRIBUTEMAPPINGMAP)
		} else {
			_, ok := v.structKeyValueMap[v.attributeMappingMap[attributeMappingMapkeyList[i].String()]]
			if !ok {
				return generateError(ERR_CODE_KEYS_MISMATCHED_IN_STRUCTKEYVALUEMAP)
			}
		}
	}
	return nil
}
