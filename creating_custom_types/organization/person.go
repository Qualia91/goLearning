package organization

import (
	"errors"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler string
}

// GetTwitterHandler returns the string twitterHandler
func (p *Person) GetTwitterHandler() string {
	return p.twitterHandler
}

// GetFirstName returns the string firstName
func (p *Person) GetFirstName() string {
	return p.firstName
}

// GetLastName returns the string lastName
func (p *Person) GetLastName() string {
	return p.lastName
}

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter handler must start with @")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) ID() string {
	return "12345"
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}
