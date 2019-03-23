package xlsToJson

import (
	"reflect"

	"github.com/tealeg/xlsx"
)

// VariableHolder - stores variables required
type VariableHolder struct {

	// currentXls - holds the current xls sheet
	currentXls *xlsx.Sheet

	// currentXlsRowCount -  row count of xls sheet
	currentXlsRowCount int

	// currentXlsColumnCount -  column count of xls sheet
	currentXlsColumnCount int

	// currentRowNo - row number of current xls sheet
	currentRowNo int

	// levelWiseNodeMap - stores the newest node at each level
	levelWiseNodeMap map[int]interface{}

	// currentNodeLevel -  current level of node
	currentNodeLevel int

	// currentNode - stores the current node
	currentNode interface{}

	// currentNodeLevelFlag - stores the current node flag
	// to be setted as APPEND_NODE_AT_NEXT_LEVEL or APPEND_NODE_AT_PREVIOUS_LEVEL or APPEND_NODE_AT_SAME_LEVEL (Defined in Constants)
	currentNodeLevelFlag string

	// dateFormatLayout - date format to be set for parsing date
	dateFormatLayout string

	// excelHeaderRow - holds the excel header row content
	excelHeaderRow *xlsx.Row

	// nodeColumnNameMap - holds the information regarding rows to be scanned for getting level
	// used to find no of levels a node can have
	nodeColumnNameMap map[string]bool

	// attributeMappingMap - holds the mapping between struct keys and excel column key
	// key :  excel column name
	// value - struct key name
	attributeMappingMap map[string]string

	// structKeyValueMap - holds the mapping between struct keys and their type
	structKeyValueMap map[string]reflect.Type

	// childrenArrayKey - used as key to append node as child
	childrenArrayKey string

	// childrenArrayEntryType - holdes the type of slice entry
	// eg : [] node => stores node
	childrenArrayEntryType reflect.Type

	// keyFunctionMap - stores the key and function to execute
	keyFunctionMap map[string]UserDefinedFunction

	// methodsInputParameterMap - map stores input to methods
	methodsInputParameterMap map[string]interface{}

	// listStrSeperator - seperator used to split the string
	listStrSeperator string

	// excelHeaderRowIndexNo - index no for excel hedaer row
	excelHeaderRowIndexNo int
}
