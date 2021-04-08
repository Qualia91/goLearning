package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LinkedList struct {
	value   int
	pointer *LinkedList
}

func main() {

	rand.Seed(time.Now().UnixNano())

	list := createList(10, nil)

	for ; list != nil; list = list.pointer {
		fmt.Printf("%v\n", list.value)
	}

	go findRandomNumber(rand.Intn(100))
	time.Sleep(5 * time.Second)

}

func findRandomNumber(randomNumber int) {

	count := 1
	numberFound := false

	for {
		number := rand.Intn(1000000000)
		if number == randomNumber {
			numberFound = true
			break
		}
		count++
	}

	if numberFound {
		fmt.Printf("Number %v found after %v attempts(s)\n", randomNumber, count)
	}

}

func createList(num int, temp *LinkedList) *LinkedList {
	if temp == nil {
		temp = &LinkedList{rand.Intn(100), nil}
		num--
	}

	tempList := temp
	for i := 0; i < num; i++ {
		t := &LinkedList{rand.Intn(100), nil}
		tempList.pointer = t
		tempList = tempList.pointer
	}

	return temp
}
