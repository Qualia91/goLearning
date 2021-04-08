package momento

// Object to store saved states
type Caretaker struct {
	history []EmployeeMomento
}

// save and revert functions in caretaker
func (c *Caretaker) Save(mo *Employee) {
	c.history = append(c.history, *mo.save())
}

func (c *Caretaker) Revert(mo *Employee) {
	// check if empty first
	if len(c.history) > 0 {
		(*mo).revert(c.history[len(c.history)-1])
		// remove last element
		c.history = c.history[:len(c.history)-1]
	}
}

// object with state to test out this pattern
type Employee struct {
	number int
	name   string
}

type EmployeeMomento struct {
	number int
	name   string
}

// add ability to save and revert to Employee
func (em *Employee) save() *EmployeeMomento {
	return NewEmployeeMomento(em.number, em.name)
}
func (em *Employee) revert(m EmployeeMomento) {
	em.number = m.number
	em.name = m.name
}

// generated
// Constructor for Caretaker
func NewCaretaker() *Caretaker {
	o := new(Caretaker)
	return o
}

// Getter method for the field history of type []Momento in the object Caretaker
func (c *Caretaker) History() []EmployeeMomento {
	return c.history
}

// Constructor for Employee
func NewEmployee(number int, name string) *Employee {
	o := new(Employee)
	o.number = number
	o.name = name
	return o
}

// Getter method for the field number of type int in the object Employee
func (e *Employee) Number() int {
	return e.number
}

// Setter method for the field number of type int in the object Employee
func (e *Employee) SetNumber(number int) {
	e.number = number
}

// Getter method for the field name of type string in the object Employee
func (e *Employee) Name() string {
	return e.name
}

// Setter method for the field name of type string in the object Employee
func (e *Employee) SetName(name string) {
	e.name = name
}

// Constructor for EmployeeMomento
func NewEmployeeMomento(number int, name string) *EmployeeMomento {
	o := new(EmployeeMomento)
	o.number = number
	o.name = name
	return o
}
