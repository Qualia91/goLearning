package main

import "fmt"

func main() {

	variadic(1, 1.2, 4.2)

	answer, _ := namedReturnVariables()

	fmt.Println(answer)

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
