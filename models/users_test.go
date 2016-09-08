package models

import (
	"testing"

	strfmt "github.com/go-openapi/strfmt"
)

func TestNewUsers(t *testing.T) {
	var reg strfmt.Registry
	users := NewUsers()

	users["bob"] = NewUser("bob", "bob@example.com")
	users["joe"] = NewUser("joe", "joe@example.com")
	users["jim"] = NewUser("jim", "jim@example.com")

	for i := range users {
		user := users[i]
		err := user.Validate(reg)
		if err != nil {
			t.Error(err)
		}
	}

	users[""] = NewUser("", "test@example.com")
	user := users[""]
	err := user.Validate(reg)
	if err == nil {
		t.Error("should fail with no name")
	}
}

func TestOutputUsers(t *testing.T) {
	users := NewUsers()

	users["bob"] = NewUser("bob", "bob@example.com")
	users["joe"] = NewUser("joe", "joe@example.com")
	users["jim"] = NewUser("jim", "jim@example.com")

	var file = "/dev/null"
	if testing.Verbose() {
		file = "/dev/stdout"
	}

	err := users.Store(file)
	if err != nil {
		t.Error(err)
	}
}
