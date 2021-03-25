package main

import (
	"deepdiveintofunctions/simplemath"
	"errors"
	"fmt"
	"io"
)

func main() {

	// this function is called at the end
	// defer functions are ordered in a stack fasion. First in, last out
	defer func() {
		println("Defer called 1")
	}()

	defer func() {
		println("Defer called 2")
	}()

	defer func() {
		if p := recover(); p == ourError {
			println("Our error occured")
			println(p)
		}
		println("Defer called 3")
	}()

	variadic(1, 1.2, 4.2)

	answer, _ := namedReturnVariables()

	fmt.Println(answer)

	sv := simplemath.NewSemanticVersion(1, 2, 3)
	println(sv.String())
	sv.IncrementMajor()
	println(sv.String())

	// annoymous function
	myFunc := func() {
		println("Hello")
	}

	for i := 0; i < 5; i++ {
		myFunc()
	}

	returnedFunc := returnFunc()

	println(returnedFunc(2, 3))

	ReadSomething()

	FunctionThatPanics()

}

var ourError = errors.New("An error occured")

func FunctionThatPanics() {
	panic(ourError)
}

func ReadSomething() error {
	// the line below gets a refernce of BadReader because the method reciever wants a pointer
	var r io.Reader = &BadReader{errors.New("error stuff")}

	if _, err := r.Read([]byte("Test something")); err != nil {
		fmt.Printf("Error occured %s\n", err)
		return err
	}

	return nil
}

type BadReader struct {
	err error
}

func (br *BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

// variadic function
func variadic(values ...float64) {
	for _, v := range values {
		fmt.Println(v)
	}
}

// named return variables
func namedReturnVariables() (answer float64, err error) {
	answer = 1.0
	return
}

// return function
func returnFunc() func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}
