package main

import "fmt"

func main() {

	switch numberToSwitchOn := 21; {
	case numberToSwitchOn > 20:
		fmt.Println("Bigger than 20")
	case numberToSwitchOn > 40:
		fmt.Println("Bigger than 40")
	}

	a := "hello"

	PrintIfString(a)
}

func PrintIfString(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("This is a string")
	default:
		fmt.Println("This is not a string")
	}
}
