package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

type europeanUnionIdNumber struct {
	country string
	value   string
}

func NewEuropeanUnionIdNumber(country string, value interface{}) Citizen {

	switch v := value.(type) {
	case string:
		return europeanUnionIdNumber{
			country: country,
			value:   v,
		}
	case int:
		return europeanUnionIdNumber{
			country: country,
			value:   strconv.Itoa(v),
		}
	case europeanUnionIdNumber:
		return europeanUnionIdNumber{
			country: country,
			value:   v.value,
		}
	default:
		panic("No type found")

	}

}

func (eui europeanUnionIdNumber) ID() string {
	return fmt.Sprintf("EU: %s\n", eui)
}

func (eui europeanUnionIdNumber) Country() string {
	return eui.country
}

type Name struct {
	firstName string
	lastName  string
}

// type alias: aliases type after = with defined type (refrences original type so cant extend)
// type TwitterHandler = string

// type decleration: Wrappes last type to make a new type (copies fields and methods to new type)
type TwitterHandler string

type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
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

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name:    Name{firstName: firstName, lastName: lastName},
		Citizen: citizen,
	}
}
