package errdefs

import "fmt"

type AppError struct {
	StatusCode int
	Message    string
	ErrorInfo  []string
}

func NewAppError(code int, msg string, args ...any) *AppError {
	var m []string
	for _, a := range args {
		switch v := a.(type) {
		case string:
			m = append(m, v)
		case []string:
			m = append(m, v...)
		default:
			m = append(m, fmt.Sprint(v))
		}
	}
	return &AppError{
		StatusCode: code,
		Message:    msg,
		ErrorInfo:  m,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

var (
	ErrEmptySpace   = &AppError{StatusCode: 400, Message: "input cannot be empty."}
	ErrInvalidValue = &AppError{StatusCode: 400, Message: "an invalid input value"}
)
