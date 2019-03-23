package xlsToJson

import (
	"log"

	"github.com/tealeg/xlsx"
)

// ReadExcel - reads excel file
func ReadExcel(excelPath string) (*xlsx.File, error) {
	xlFile, err := xlsx.OpenFile(excelPath)
	if err != nil {
		log.Fatal("ReadExcel : error in opening excel : ", err)
		return nil, generateError(ERR_CODE_EXCEL_NOT_OPEN)
	}
	return xlFile, nil
}
