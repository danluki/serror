package main

import (
	"fmt"

	"github.com/danluki/serror"
	"github.com/danluki/serror/examples/simple_service/errors"
)

func GetHttpError(err error) error {
	switch err := err.(type) {
	case errors.ObjectBadState:
		return fmt.Errorf("Object bad state %s", err.Message)
	case errors.BadRequestError:
		return fmt.Errorf("Bad request %s", err.Message)
	default:
		return fmt.Errorf("Internal server error")
	}
}

func testHandler() error {
	var err error
	err = errors.NewObjectBadState("test", "test", "test")
	e := GetHttpError(err)
	if e != nil {
		fmt.Println(e)
	}

	marshalledErr, err := serror.Marhal(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(marshalledErr))
	err = errors.Unmarshal(marshalledErr)
	fmt.Println(err)

	return nil
}

func main() {
	err := testHandler()
	if err != nil {
		fmt.Println(err)
	}

	// err = errors.NewNotFoundError("", map[string]string{
	// 	"instance": "user",
	// 	"id":       "1",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
