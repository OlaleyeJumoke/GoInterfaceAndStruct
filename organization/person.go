package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//Identifiable interface
type Identifiable interface {
	ID() string
}

//Citizen interface
type Citizen interface {
	Country() string
	Identifiable
}

type socialSecurityNumber string

//NewSocialSecurityNumber value
func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

//ID string
func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

//ID string
func (ssn socialSecurityNumber) Country() string {
	return string(ssn)
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

//NewEuropeanUnionIdentifier value
// making the id an interface makes it a dynamic type(which could be anytype)
func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: country,
		}
	case europeanUnionIdentifier:
		return v
	default:
		panic("Wrong type")
	}

	//return europeanUnionIdentifier{
	//	id:      id,
	//	country: country,
	//}
}

//ID string
func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

//ID string
func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("Eu: %s", eui.country)
}

// Name strct
type Name struct {
	first string
	last  string
}

//FullName hj
func (n Name) FullName() string {
	return fmt.Sprintf("Format: %s %s", n.first, n.last)
}

//Employee data
type Employee struct {
	Name
}

//Handler ok
type Handler struct {
	handler string
	name    string
}

// TwitterHandler kk
type TwitterHandler string

//RedirectURL redirects url
func (th TwitterHandler) RedirectURL() string {

	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://twitter.com/%s", cleanHandler)
}

//Person struct
type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}

// NewPerson person
func NewPerson(lastName, firstName string, identifiable Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: identifiable,
	}

}

//SetTwitterHandler jky
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("An handler must start with @")
	}
	p.twitterHandler = handler
	return nil
}

//TwitterHandler prints the string
func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
