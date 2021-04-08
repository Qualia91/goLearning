package state

type State interface {
	HandleRequest()
	DebugName() string
}

type State1 struct {
	fan       *Fan
	debugName string
}
type State2 struct {
	fan       *Fan
	debugName string
}
type State3 struct {
	fan       *Fan
	debugName string
}

// Implements State
func (s *State1) HandleRequest() {
	s.fan.SetCurrentState(s.fan.State2())
}

// Implements State
func (s *State2) HandleRequest() {
	s.fan.SetCurrentState(s.fan.State3())
}

// Implements State
func (s *State3) HandleRequest() {
	s.fan.SetCurrentState(s.fan.State1())
}

// Constructor for State1
func NewState1(fan *Fan) *State1 {
	o := new(State1)
	o.fan = fan
	o.debugName = "state1"
	return o
}

// Getter method for the field fan of type *Fan in the object State1
func (s *State1) Fan() *Fan {
	return s.fan
}

// Constructor for State2
func NewState2(fan *Fan) *State2 {
	o := new(State2)
	o.fan = fan
	o.debugName = "state2"
	return o
}

// Getter method for the field fan of type *Fan in the object State2
func (s *State2) Fan() *Fan {
	return s.fan
}

// Constructor for State3
func NewState3(fan *Fan) *State3 {
	o := new(State3)
	o.fan = fan
	o.debugName = "state3"
	return o
}

// Getter method for the field fan of type *Fan in the object State3
func (s *State3) Fan() *Fan {
	return s.fan
}

// Getter method for the field debugName of type string in the object State1
func (s *State1) DebugName() string {
	return s.debugName
}

// Getter method for the field debugName of type string in the object State2
func (s *State2) DebugName() string {
	return s.debugName
}

// Getter method for the field debugName of type string in the object State3
func (s *State3) DebugName() string {
	return s.debugName
}
