package errors

// SError represents the internal error structure
type SError struct {
	Code          int
	Description   string
	InternalError error
}

// NewCurrencyError initializes the SError
func NewCurrencyError(code int, description string, err error) *SError {
	sErr := &SError{
		InternalError: err,
		Code:          code,
		Description:   description,
	}

	return sErr
}

func (e *SError) Error() string {
	return e.Description
}
