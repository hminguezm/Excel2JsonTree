package xlsToJson

// UserDefinedFunction - user defined functionality  type
type UserDefinedFunction func(map[string]interface{}, string, *VariableHolder) (interface{}, error)
