package main

import "fmt"

func main() {

	list := []int{1, 4, 7, 11, 15, 32, 67, 99}

	fmt.Printf("%t", BTS(list, 1))

}

func BTS(list []int, searchNumber int) bool {

	// get length of list
	listLength := len(list) / 2

	// check list length is over 1
	if listLength <= 0 {
		return list[listLength] == searchNumber
	} else {
		// get value at this index
		if list[listLength] == searchNumber {
			return true
		} else if list[listLength] < searchNumber {
			return BTS(list[listLength+1:], searchNumber)
		} else {
			return BTS(list[:listLength], searchNumber)
		}
	}

	return false

}
