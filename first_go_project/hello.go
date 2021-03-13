package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	path := flag.String("path", "myapp.log", "The path to log")
	level := flag.String("level", "ERROR", "The level")

	// looks at input flags
	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	// defer - happens after we are done with it
	defer f.Close()

	r := bufio.NewReader(f)

	// infinite for loop
	for {
		s, err := r.ReadString('\n')
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
		if err != nil {
			break
		}
	}
}
