package state

type Fan struct {
	state1       State
	state2       State
	state3       State
	currentState State
}

func (f *Fan) PullChain() {
	f.currentState.HandleRequest()
}

// Constructor for Fan
func NewFan() *Fan {
	o := new(Fan)
	o.state1 = NewState1(o)
	o.state2 = NewState2(o)
	o.state3 = NewState3(o)
	o.currentState = o.state1
	return o
}

// Getter method for the field state1 of type State in the object Fan
func (f *Fan) State1() State {
	return f.state1
}

// Setter method for the field state1 of type State in the object Fan
func (f *Fan) SetState1(state1 State) {
	f.state1 = state1
}

// Getter method for the field state2 of type State in the object Fan
func (f *Fan) State2() State {
	return f.state2
}

// Setter method for the field state2 of type State in the object Fan
func (f *Fan) SetState2(state2 State) {
	f.state2 = state2
}

// Getter method for the field state3 of type State in the object Fan
func (f *Fan) State3() State {
	return f.state3
}

// Setter method for the field state3 of type State in the object Fan
func (f *Fan) SetState3(state3 State) {
	f.state3 = state3
}

// Getter method for the field currentState of type State in the object Fan
func (f *Fan) CurrentState() State {
	return f.currentState
}

// Setter method for the field currentState of type State in the object Fan
func (f *Fan) SetCurrentState(currentState State) {
	f.currentState = currentState
}
