package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	stingChan := make(chan string)
	tower1Chan := make(chan string)
	tower2Chan := make(chan string)

	var offset int32 = 3

	go tower1(stingChan, tower1Chan, offset)
	go tower2(stingChan, tower2Chan, offset)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-tower1Chan:
			fmt.Printf("\nControl Tower: Message from tower 1 - %v", msg)
		case msg := <-tower2Chan:
			fmt.Printf("\nControl Tower: Message from tower 2 - %v", msg)
		}
	}
}

func tower1(s chan string, t1 chan string, offset int32) {

	inputStream := bufio.NewReader(os.Stdin)
	fmt.Println("Tower 1: Enter message: ")
	userInput, _ := inputStream.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)

	var secretString string

	for _, c := range userInput {
		secretString += string(c + offset)
	}

	s <- secretString
	t1 <- "Message Sent to 2"
}

func tower2(s chan string, t2 chan string, offset int32) {

	secretString := <-s

	var orgString string

	for _, c := range secretString {
		orgString += string(c - offset)
	}

	fmt.Printf("%v", orgString)

	t2 <- "Message Recieved from 1"
}
