package errors

type ErrorDetail struct {
	Message string
	Code    int
}

func NewErrorDetail(message string, code int) *ErrorDetail {
	return &ErrorDetail{
		Message: message,
		Code:    code,
	}
}

func (e *ErrorDetail) Error() string {
	return e.Message
}
