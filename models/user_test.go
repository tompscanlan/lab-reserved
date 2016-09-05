package models

import (
	"log"
	"testing"

	strfmt "github.com/go-openapi/strfmt"
)

func TestNewUser(t *testing.T) {
	var reg strfmt.Registry

	user := NewUser("name", "tom@test.com")

	log.Println(user)
	err := user.Validate(reg)
	if err != nil {
		t.Error(err)
	}

	user = NewUser("name", "xxx")
	err = user.Validate(reg)
	if err != nil {
		t.Error(err)
	}

}
