package http

import (
	"github.com/danielalmeidafarias/go-clean/internal/errors"
)

var HttpStatusCode map[errors.ErrorStatus]int = map[errors.ErrorStatus]int{
	errors.BadRequest:         400,
	errors.NotFound:           404,
	errors.Conflict:           409,
	errors.FailedPrecondition: 412,
	errors.Internal:           500,
}

func ValidationError(err error) map[string]string {
	return map[string]string{
		"validation error": err.Error(),
	}
}
