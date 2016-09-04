package models

import (
	"log"
	"testing"

	strfmt "github.com/go-openapi/strfmt"
)

func TestNewOwner(t *testing.T) {
	var reg strfmt.Registry

	owner := NewOwner("name", "tom@test.com")

	log.Println(owner)
	err := owner.Validate(reg)
	if err != nil {
		t.Error(err)
	}

	owner = NewOwner("name", "xxx")
	err = owner.Validate(reg)
	if err != nil {
		t.Error(err)
	}

}
