package errors

import (
	"fmt"

	"github.com/danluki/serror"
)

const (
	ErrObjectBadState = "ObjectBadState"
)

type ObjectBadState struct {
	serror.Base

	ObjectType string
	ObjectId   string
	State      string
}

func NewObjectBadState(objectType, objectId, state string) ObjectBadState {
	msg := fmt.Sprintf("%s in bad state: %s", objectType, state)

	return ObjectBadState{
		Base: serror.Base{
			Message: msg,
			Type:    ErrObjectBadState,
		},
		ObjectType: objectType,
		ObjectId:   objectId,
		State:      state,
	}
}
