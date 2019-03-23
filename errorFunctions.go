package xlsToJson

import "errors"

// generateError - accepts the string and returns error
func generateError(errorString string) error {
	return errors.New(errorString)
}
