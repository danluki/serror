package errors

import (
	"fmt"
	"strings"

	"github.com/danluki/serror"
)

const (
	BadRequestErrorType = "BadRequest"
)

type BadRequestError struct {
	serror.Base

	Reason   []string
	HttpCode int
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("This is custom error message for test: %s", e.Message)
}

func NewBadRequestError(reason []string) BadRequestError {
	msg := "Bad request: " + strings.Join(reason, ", ")
	return BadRequestError{
		Base: serror.Base{
			Message: msg,
			Type:    BadRequestErrorType,
		},
		Reason:   reason,
		HttpCode: 400,
	}
}
