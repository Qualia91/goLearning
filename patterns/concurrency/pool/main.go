package main

import (
	"fmt"
	"sync"
)

func main() {

	// create a pool of objects. Pools have one field = a function which creates a new object
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}

	// get an object. No objects are in the pool so one is created
	myPool.Get()

	// create another object
	inst := myPool.Get()

	// now we are finished with inst object, put it back into the pool
	myPool.Put(inst)

	// now when we get an object, we get the one we just put in
	myPool.Get()

}
