package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(4) + 1

		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("Found in cache")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("Found in database")
				fmt.Println(b)
			}
		}(id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (b Book, ok bool) {
	b, ok = cache[id]
	return
}

func queryDatabase(id int) (b Book, ok bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b = range books {
		if ok = b.ID == id; ok {
			cache[id] = b
			return
		}
	}

	ok = false
	return
}
