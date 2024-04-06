package validations

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

//func GenerateValidationErrors(err error) *[]ValidationError {
//var validationErrors []ValidationError
//var validationError echo.Validator
//}
