package xlsToJson

const (

	//EMPTY_STR - empty string
	EMPTY_STR = ""

	// DEFAULT_LIST_STRING_SEPERATOR - default seperator
	DEFAULT_LIST_STRING_SEPERATOR = ","

	//DEFAULT_DATE_FORMAT - default layout to parse time value
	DEFAULT_DATE_FORMAT_LAYOUT = "01-02-2006"

	//DEFAULT_HEADER_ROW_INDEX	- default index of header row
	DEFAULT_HEADER_ROW_INDEX = 0

	//OPEN_CLOSE_SQUARE_BRACKET - constant for reflect type squrare bracket opening and closing square bracket
	OPEN_CLOSE_SQUARE_BRACKET = "[]"

	//STRING - constant for reflect type string
	STRING = "string"

	//INT - constant for reflect type int
	INT = "int"

	//INT8 - constant for reflect type int8
	INT8 = "int8"

	//INT32 - constant for reflect type int32
	INT32 = "int32"

	//INT64 - constant for reflect type int64
	INT64 = "int64"

	//FLOAT32 - constant for reflect type float32
	FLOAT32 = "float32"

	//FLOAT64 - constant for reflect type float64
	FLOAT64 = "float64"

	//BOOLEAN - constant for reflect type bool
	BOOLEAN = "bool"

	//STRUCT - constant for reflect Kind bool
	STRUCT = "struct"

	//SLICE - constant for reflect Kind slice
	SLICE = "slice"

	// TIME_DOT_TIME - constanst for reflect type time.time
	TIME_DOT_TIME = "time.Time"

	//INT_SLICE - constant for reflect type slice of int ([]int)
	INT_SLICE = OPEN_CLOSE_SQUARE_BRACKET + INT

	//INT8_SLICE - constant for reflect type slice of int8 ([]int8)
	INT8_SLICE = OPEN_CLOSE_SQUARE_BRACKET + INT8

	//INT32_SLICE - constant for reflect type slice of int32 ([]int32)
	INT32_SLICE = OPEN_CLOSE_SQUARE_BRACKET + INT32

	//INT64_SLICE - constant for reflect type slice of int64 ([]int64)
	INT64_SLICE = OPEN_CLOSE_SQUARE_BRACKET + INT64

	//FLOAT_32SLICE - constant for reflect type slice of float32 ([]float32)
	FLOAT32_SLICE = OPEN_CLOSE_SQUARE_BRACKET + FLOAT32

	//INT64_SLICE - constant for reflect type slice of float64 ([]float64)
	FLOAT64_SLICE = OPEN_CLOSE_SQUARE_BRACKET + FLOAT64

	//STRING_SLICE - constant for reflect type slice of string ([]string)
	STRING_SLICE = OPEN_CLOSE_SQUARE_BRACKET + STRING

	//BOOLEAN_SLICE - constant for reflect type slice of bool ([]bool)
	BOOLEAN_SLICE = OPEN_CLOSE_SQUARE_BRACKET + BOOLEAN

	// NODE FLAGS - indicate level of flag
	APPEND_NODE_AT_START_LEVEL    = "APPEND_NODE_AT_START_LEVEL"
	APPEND_NODE_AT_NEXT_LEVEL     = "APPEND_NODE_AT_NEXT_LEVEL"
	APPEND_NODE_AT_PREVIOUS_LEVEL = "APPEND_NODE_AT_PREVIOUS_LEVEL"
	APPEND_NODE_AT_SAME_LEVEL     = "APPEND_NODE_AT_SAME_LEVEL"

	// EXCEL_TAG - tag to be used while mapping excel column key with struct key
	EXCEL_TAG = "excel"
)
