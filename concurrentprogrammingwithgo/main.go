package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(4) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("Found in cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("Found in database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		//time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (b Book, ok bool) {
	m.RLock()
	b, ok = cache[id]
	m.RUnlock()
	return
}

func queryDatabase(id int, m *sync.RWMutex) (b Book, ok bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b = range books {
		if ok = b.ID == id; ok {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return
		}
	}

	ok = false
	return
}
