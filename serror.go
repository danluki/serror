package serror

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Base struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (b Base) Error() string {
	return fmt.Sprintf("%s: %s", b.Type, b.Message)
}

func Marhal(
	inerr interface{},
) ([]byte, error) {
	errBytes, err := json.Marshal(inerr)
	if err != nil {
		return nil, err
	}

	return errBytes, nil
}

func Unmarshal(
	body []byte,
	errInstances map[string]interface{},
) error {
	var base Base

	err := json.Unmarshal(body, &base)
	if err != nil {
		return err
	}

	var instance interface{}

	if _, ok := errInstances[base.Type]; ok {
		instance = errInstances[base.Type]
	} else {
		return fmt.Errorf("Unknown error type: %s", base.Type)
	}

	target := reflect.New(reflect.TypeOf(instance))
	err = json.Unmarshal(body, target.Interface())
	if err != nil {
		return err
	}

	err = target.Elem().Interface().(error)
	return err
}
