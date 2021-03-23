package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

// type alias: aliases type after = with defined type (refrences original type so cant extend)
// type TwitterHandler = string

// type decleration: Wrappes last type to make a new type (copies fields and methods to new type)
type TwitterHandler string

type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

func (th TwitterHandler) RedirectUrl() string {
	cleanhandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanhandler)
}

// GetTwitterHandler returns the string twitterHandler
func (p *Person) GetTwitterHandler() TwitterHandler {
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

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with @")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) ID() TwitterHandler {
	return "12345"
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}
