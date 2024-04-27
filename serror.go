package serror

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Base struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (b Base) Error() string {
	return fmt.Sprintf("%s: %s", b.Type, b.Message)
}

func ParseArgs(msg string, args map[string]string) (string, error) {
	parsed := msg
	for k, v := range args {
		parsed = strings.ReplaceAll(parsed, "{{"+k+"}}", v)
	}

	if strings.Contains(parsed, "{{") {
		return "", fmt.Errorf(
			"message '%s' contains unreplaced placeholders, parsing result: '%s'",
			msg,
			parsed,
		)
	}

	return parsed, nil
}

func MustParseArgs(msg string, args map[string]string) string {
	parsed, err := ParseArgs(msg, args)
	if err != nil {
		panic(err)
	}

	return parsed
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
