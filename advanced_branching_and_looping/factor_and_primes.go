package main

import "fmt"

func main() {

	numbers := []int{1, 2, 4, 8, 47, 94, 188, 376}

	for _, currentNumber := range numbers {

		for i := 1; i < currentNumber; i++ {

			if currentNumber%i == 0 {
				fmt.Printf("%v is a factor of %v\n", i, currentNumber)
			}

		}

	}

}
