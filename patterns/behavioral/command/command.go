package command

// The command interface, used to create objects that contain the information necessary to executer and undo
type Command struct {
	execute   func()
	unExecute func()
}

// Constructor for Command
func NewCommand(execute func(), unExecute func()) *Command {
	o := new(Command)
	o.execute = execute
	o.unExecute = unExecute
	return o
}

// Getter method for the field execute of type func() in the object Command
func (c *Command) Execute() func() {
	return c.execute
}

// Setter method for the field execute of type func() in the object Command
func (c *Command) SetExecute(execute func()) {
	c.execute = execute
}

// Getter method for the field unExecute of type func() in the object Command
func (c *Command) UnExecute() func() {
	return c.unExecute
}

// Setter method for the field unExecute of type func() in the object Command
func (c *Command) SetUnExecute(unExecute func()) {
	c.unExecute = unExecute
}
