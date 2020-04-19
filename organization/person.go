package organization

import (
	"errors"
	"fmt"
	"strings"
)

//Identifiable interface
type Identifiable interface {
	ID() string
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
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

// NewPerson person
func NewPerson(lastName, firstName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName}
}

//FullName hj
func (p *Person) FullName() string {
	return fmt.Sprintf("Format: %s %s", p.firstName, p.lastName)
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

//ID string
func (p *Person) ID() string {
	return "12334"
}
