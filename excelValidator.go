package xlsToJson

import "github.com/tealeg/xlsx"

// isRowLengthZero - returns true if sheet row list length is zero
func isRowLengthZero(sheetRowList []*xlsx.Row) bool {
	if len(sheetRowList) == 0 {
		return true
	}
	return false
}

// isColumnLengthZero - returns true if sheet column list length is zero
func isColumnLengthZero(sheetColumnList []*xlsx.Col) bool {
	if len(sheetColumnList) == 0 {
		return true
	}
	return false
}
