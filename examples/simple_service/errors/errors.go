package errors

import (
	"github.com/danluki/serror"
)

var errInstances = map[string]interface{}{
	ErrObjectBadState:   ObjectBadState{},
	BadRequestErrorType: BadRequestError{},
}

func Unmarshal(body []byte) error {
	return serror.Unmarshal(body, errInstances)
}
