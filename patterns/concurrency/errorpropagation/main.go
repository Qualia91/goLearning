package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

// custom error type to include everything that should be in error
type CustomError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func WrapError(inner error, message string, msgArgs ...interface{}) CustomError {
	return CustomError{
		Inner:      inner,
		Message:    fmt.Sprintf(message, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

func main() {

	var ce interface{}
	ce = WrapError(errors.New("Hello"), "World")

	if err, ok := ce.(error); ok {
		fmt.Println(err.Error())
	} else if cet, ok := ce.(CustomError); ok {
		fmt.Println(cet.Error())
	} else {
		fmt.Println("Unknown error type")
	}
}
