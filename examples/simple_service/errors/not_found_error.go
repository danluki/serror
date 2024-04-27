package errors

import (
	"net/http"

	"github.com/danluki/serror"
)

const (
	NotFoundErrorType = "NotFoundError"
)

type NotFoundError struct {
	serror.Base

	Code int
}

func NewNotFoundError(msg string, args map[string]string) NotFoundError {
	var parsedMsg string

	if msg == "" {
		msg = "{{instance}}: {{id}} not found"
	}
	parsedMsg = serror.MustParseArgs(msg, args)

	return NotFoundError{
		Base: serror.Base{
			Message: parsedMsg,
			Type:    NotFoundErrorType,
		},
		Code: http.StatusNotFound,
	}
}
