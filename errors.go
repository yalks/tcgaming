package tcgaming

import "fmt"

type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("TC Gaming API error: %s (code: %d)", e.Message, e.Code)
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func IsAPIError(err error) bool {
	_, ok := err.(*APIError)
	return ok
}

func GetAPIErrorCode(err error) (int, bool) {
	apiErr, ok := err.(*APIError)
	if !ok {
		return 0, false
	}
	return apiErr.Code, true
}