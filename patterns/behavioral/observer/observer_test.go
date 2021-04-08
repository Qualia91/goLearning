package observer_test

import (
	"behavioral/observer"
	"testing"
)

func TestObserver(t *testing.T) {
	observableImpl := observer.NewObservableImpl()

	observer1 := observer.NewObserverImpl(1, "Hello")
	observer2 := observer.NewObserverImpl(2, "World")

	observableImpl.Add(observer1)
	observableImpl.Add(observer2)

	observableImpl.Notify("Hello")
	observableImpl.Notify("World")

	observableImpl.Remove(observer1)

	observableImpl.Notify("Hello")

	if observer1.LastMessage() != "World" {
		t.Errorf("Last message observer1 received is wrong. It was %v, and it should have been %v\n", observer1.LastMessage(), "World")
	}

	if observer2.LastMessage() != "Hello" {
		t.Errorf("Last message observer2 received is wrong. It was %v, and it should have been %v\n", observer2.LastMessage(), "Hello")
	}

}

func Example() {
	// Create object to be observed
	observableImpl := observer.NewObservableImpl()

	// create observers
	observer1 := observer.NewObserverImpl(1, "Hello")
	observer2 := observer.NewObserverImpl(2, "World")

	// add observers
	observableImpl.Add(observer1)
	observableImpl.Add(observer2)

	// change state in observable which will update observers
	observableImpl.Notify("Hello")
	observableImpl.Notify("World")

	// remove an observer
	observableImpl.Remove(observer1)

	// update observable with value that remaining observer wont observe
	observableImpl.Notify("Hello")

	//Output:Observable with id 1 has accepted state of Hello
	//Observable with id 2 has accepted state of World
}
