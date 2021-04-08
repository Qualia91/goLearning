package observer

import (
	"fmt"
	"sync"
)

type Observable interface {
	Add(observer Observer)
	Remove(observer Observer)
	Notify(event interface{})
}

type Observer interface {
	NotifyCallback(event interface{})
}

type ObservableImpl struct {
	observer sync.Map
}

type ObserverImpl struct {
	id          int
	state       string
	lastMessage string
}

// Implements Observable
func (observableImpl *ObservableImpl) Add(observer Observer) {
	observableImpl.observer.Store(observer, struct{}{})
}

// Implements Observable
func (observableImpl *ObservableImpl) Remove(observer Observer) {
	observableImpl.observer.Delete(observer)
}

// Implements Observable
func (observableImpl *ObservableImpl) Notify(event interface{}) {
	observableImpl.observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}

		// type assertion on Observer Type
		key.(Observer).NotifyCallback(event)
		return true
	})
}

// Implements Observer
func (observerImpl *ObserverImpl) NotifyCallback(event interface{}) {
	if event.(string) == observerImpl.state {
		fmt.Printf("Observable with id %d has accepted state of %s\n", observerImpl.id, event)
	}
	observerImpl.lastMessage = event.(string)
}

// Constructor for ObservableImpl
func NewObservableImpl() *ObservableImpl {
	o := new(ObservableImpl)
	return o
}

// Constructor for ObserverImpl
func NewObserverImpl(id int, state string) *ObserverImpl {
	o := new(ObserverImpl)
	o.id = id
	o.state = state
	return o
}

// Getter method for the field lastMessage of type string in the object ObserverImpl
func (o *ObserverImpl) LastMessage() string {
	return o.lastMessage
}
