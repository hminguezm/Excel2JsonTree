package xlsToJson

import (
	"github.com/tealeg/xlsx"
)

// checkAndSetCurrentXlsInfo - check and set the current xls info
func (v *VariableHolder) checkAndSetCurrentXlsInfo(sheet *xlsx.Sheet, childrenArrayKey string) error {
	if isColumnLengthZero(sheet.Cols) {
		return generateError(ERR_CODE_EXCEL_SHEET_EMPTY)
	}
	if isRowLengthZero(sheet.Rows) {
		return generateError(ERR_CODE_EXCEL_SHEET_EMPTY)
	}
	v.currentXls = sheet
	if len(v.currentXls.Rows[v.excelHeaderRowIndexNo].Cells) == 0 {
		return generateError(ERR_CODE_SHEET_HEADER_ROW_EMPTY)
	}
	v.excelHeaderRow = v.currentXls.Rows[v.excelHeaderRowIndexNo]
	v.currentXlsColumnCount = len(v.excelHeaderRow.Cells)
	v.currentXlsRowCount = len(v.currentXls.Rows)
	return nil
}