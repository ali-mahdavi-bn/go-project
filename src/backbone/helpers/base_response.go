package helpers

import "go-tel/src/backbone/helpers/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       int                            `json:"result_code"`
	ValidationErrors *[]validations.ValidationError `json:"validation_errors"`
	Error            any                            `json:"error"`
}
