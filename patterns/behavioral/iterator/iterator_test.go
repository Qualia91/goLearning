package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {

	myCollection := NewMyCollection()

	myCollection.Add(1)
	myCollection.Add(2)
	myCollection.Add(3)
	myCollection.Add(4)
	myCollection.Add(5)
	myCollection.Add(6)
	myCollection.Add(7)
	myCollection.Add(8)

	iter := myCollection.IterC()

	counter := 1
	for {
		if !iter.HasNext() {
			break
		}
		next := iter.Next()
		if next != counter {
			t.Errorf("Iter test failed, Expecting %v, received %v", counter, next)
		}
		counter++
	}

	if counter != 9 {
		t.Errorf("Iter test failed, expected it to run %v times, ran %v times", 8, counter-1)
	}

}

func BenchmarkIterator(b *testing.B) {

	collectionSize := 1000000

	myCollection := NewMyCollection()

	for i := 0; i < collectionSize; i++ {
		myCollection.Add(i)
	}

	iter := myCollection.IterC()

	counter := 0
	for {
		if !iter.HasNext() {
			break
		}
		next := iter.Next()
		if next != counter {
			b.Errorf("Iter benchmark failed, Expecting %v, received %v", counter, next)
		}
		counter++
	}

	if counter != collectionSize {
		b.Errorf("Iter benchmark failed, expected it to run %v times, ran %v times", collectionSize, counter)
	}

}

func Example() {

	// create collection
	myCollection := NewMyCollection()

	// add to collection
	myCollection.Add(1)
	myCollection.Add(2)
	myCollection.Add(3)
	myCollection.Add(4)
	myCollection.Add(5)
	myCollection.Add(6)
	myCollection.Add(7)
	myCollection.Add(8)

	// create iterator
	iter := myCollection.IterC()

	// iterate over while has next
	for {
		if !iter.HasNext() {
			break
		}
		fmt.Printf("Next object in MyCollection is %v\n", iter.Next())

	}

	//Output:Next object in MyCollection is 1
	//Next object in MyCollection is 2
	//Next object in MyCollection is 3
	//Next object in MyCollection is 4
	//Next object in MyCollection is 5
	//Next object in MyCollection is 6
	//Next object in MyCollection is 7
	//Next object in MyCollection is 8
}
