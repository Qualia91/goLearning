package command

// Simple object to be the receiver of commands
type Receiver struct {
	On bool
}

func NewReceiver() *Receiver {
	r := new(Receiver)
	r.On = false
	return r
}
