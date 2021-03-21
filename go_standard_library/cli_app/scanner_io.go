package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var age int
		var name string

		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)
		if err != nil {
			panic(err)
		}

		if n == 2 {
			fmt.Printf("%s, %d\n", name, age)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
