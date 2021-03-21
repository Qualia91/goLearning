package mylib

import "fmt"

func adder(l, r int) int {
	return l + r
}

func subtractor(l, r int) int {
	return l - r
}

func messageWriter(greeting, name string) string {
	message := fmt.Sprintf("%v, %v", greeting, name)
	return message
}
