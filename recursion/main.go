package main

import "fmt"

func main() {
	fmt.Println(<-Fibonacci(10))
}

func Fibonacci(n int) <-chan int {
	results := make(chan int)

	go func(n int) {
		if n <= 2 {
			results <- 1
			return
		}
		results <- <-Fibonacci(n-1) + <-Fibonacci(n-2)
	}(n)

	return results
}
