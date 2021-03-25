package main

import (
	"deepdiveintofunctions/simplemath"
	"fmt"
)

func main() {

	variadic(1, 1.2, 4.2)

	answer, _ := namedReturnVariables()

	fmt.Println(answer)

	sv := simplemath.NewSemanticVersion(1, 2, 3)
	println(sv.String())
	sv.IncrementMajor()
	println(sv.String())

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
