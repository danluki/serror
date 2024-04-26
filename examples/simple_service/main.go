package main

import (
	"fmt"

	"github.com/danluki/serror"
	"github.com/danluki/serror/examples/simple_service/errors"
)

func GetHttpError(err error) error {
	switch err := err.(type) {
	case errors.ObjectBadState:
		return fmt.Errorf("Object bad state", err.Message)
	case errors.BadRequestError:
		return fmt.Errorf("Bad request", err.Message)
	default:
		return fmt.Errorf("Internal server error")
	}
}

func testHandler() error {
	var err error
	err = errors.NewObjectBadState("test", "test", "test")

	marshalledErr, err := serror.Marhal(err)
	if err != nil {
		return err
	}

	fmt.Println(string(marshalledErr))
	err = errors.Unmarshal(marshalledErr)
	fmt.Println(err)
	e := GetHttpError(err)
	if e != nil {
		return e
	}

	return nil
}

func main() {
	err := testHandler()
	if err != nil {
		fmt.Println(err)
	}
}
