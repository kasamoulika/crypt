package models

import "crypt.com/v2/application/errors"

// ErrorResponse represents error response structure
type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func GetErrorResponse(sError *errors.SError) *ErrorResponse {
	return &ErrorResponse{
		Code:        sError.Code,
		Description: sError.Description,
	}
}
