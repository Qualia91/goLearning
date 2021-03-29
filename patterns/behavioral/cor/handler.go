package cor

import "fmt"

type Handler interface {
	SetSuccessor(Handler)
	HandleRequest(Request)
}

type Director struct {
	successor Handler
}

// Implements Handler
func (dir *Director) SetSuccessor(handler Handler) {
	dir.successor = handler
}

// Implements Handler
func (dir *Director) HandleRequest(request Request) {
	if request.amount < 100 {
		fmt.Printf("Amount of %v is less than 100 so director can action it\n", request.amount)
		return
	}
	dir.successor.HandleRequest(request)
}

type VP struct {
	successor Handler
}

// Implements Handler
func (vp *VP) SetSuccessor(handler Handler) {
	vp.successor = handler
}

// Implements Handler
func (vp *VP) HandleRequest(request Request) {
	fmt.Printf("Amount of %v is more than 100 so VP can action it\n", request.amount)
}

// Constructor for Director
func NewDirector() *Director {
	o := new(Director)
	return o
}

// Constructor for VP
func NewVP() *VP {
	o := new(VP)
	return o
}
