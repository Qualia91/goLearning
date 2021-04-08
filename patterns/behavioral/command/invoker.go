package command

import "fmt"

type Invoker struct {
	commandStack []Command
}

func (i *Invoker) Execute(command Command) {
	command.execute()
	i.commandStack = append(i.commandStack, command)
}

func (i *Invoker) Undo() {
	if len(i.commandStack) == 0 {
		fmt.Println("Command Stack is Empty")
	} else {
		i.commandStack[len(i.commandStack)-1].unExecute()
		i.commandStack = i.commandStack[:len(i.commandStack)-1]
	}
}

// Getter method for the field commandStack of type []Command in the object Invoker
func (i *Invoker) CommandStack() []Command {
	return i.commandStack
}

// Constructor for Invoker
func NewInvoker() *Invoker {
	o := new(Invoker)
	return o
}
