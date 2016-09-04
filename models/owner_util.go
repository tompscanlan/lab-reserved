package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//	"github.com/asaskevich/govalidator"
	"github.com/go-openapi/strfmt"
)

func NewOwner(name string, email string) *Owner {
	// at some point, need to add in email validation
	//	if !govalidator.IsEmail(email) {
	//		log.Printf("Got bad email (%s) for owner (%s)", email, name)
	//		email = "bad@example.com"
	//	}

	e := strfmt.Email(email)

	owner := new(Owner)
	owner.Name = &name
	owner.Email = &e
	return owner
}

func (owner Owner) String() string {

	return fmt.Sprintf("owner: \"%s\" <%s>", *owner.Name, owner.Email)

}

// Store all the items as JSON in a file
func (owner Owner) Store(filename string) error {
	b, err := json.Marshal(owner)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
