package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//	"github.com/asaskevich/govalidator"
	"github.com/go-openapi/strfmt"
)

func NewUser(name string, email string) User {
	// at some point, need to add in email validation
	//	if !govalidator.IsEmail(email) {
	//		log.Printf("Got bad email (%s) for user (%s)", email, name)
	//		email = "bad@example.com"
	//	}

	e := strfmt.Email(email)

	user := new(User)
	user.Name = &name
	user.Email = &e
	return *user
}

func (user User) String() string {

	return fmt.Sprintf("user: \"%s\" <%s>", *user.Name, user.Email)

}

// Store all the items as JSON in a file
func (user User) Store(filename string) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
