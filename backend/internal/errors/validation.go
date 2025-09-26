package errdefs

import "fmt"

type ValueErr struct {
	StatusCode int
	Message    string
	ErrorInfo  []string
}

func NewAppError(code int, msg string, args ...any) *ValueErr {
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
	return &ValueErr{
		StatusCode: code,
		Message:    msg,
		ErrorInfo:  m,
	}
}

func (e *ValueErr) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

var (
	ErrEmptySpace   = &ValueErr{StatusCode: 400, Message: "input cannot be empty."}
	ErrInvalidValue = &ValueErr{StatusCode: 400, Message: "an invalid input value"}
)
