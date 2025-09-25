package errdefs

import "fmt"

type AppError struct {
	StatusCode int
	Message    string
	ErrorInfo  interface{}
}

func NewAppError(code int, msg string, args ...any) *AppError {
	var m interface{}
	if len(args) > 0 {
		if len(args) == 1 {
			m = args[0]
		} else {
			m = args
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
	ErrInvalidValue = NewAppError(400, "an invalid input value")
	ErrEmptySpace   = NewAppError(400, "an input value should not have an empty space")
)
