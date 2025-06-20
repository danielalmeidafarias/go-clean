package errors

import "fmt"

type ErrorStatus int

const (
	NotFound ErrorStatus = iota
	BadRequest
	Conflict
	FailedPrecondition
	Internal
)

type Error struct {
	Message string
	Code    ErrorStatus
}

func NewError(message string, code ErrorStatus) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func FromError(err error, code ErrorStatus) *Error {
	return &Error{
		Message: err.Error(),
		Code:    code,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) WithContext(context string) *Error {
	return &Error{
		Message: fmt.Sprintf("%s: %s", context, e.Message),
		Code:    e.Code,
	}
}
