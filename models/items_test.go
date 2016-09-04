package models

import (
	"testing"

	strfmt "github.com/go-openapi/strfmt"
)

func TestNewItems(t *testing.T) {
	var reg strfmt.Registry
	items := NewItems()

	items = append(items, NewItem("name", "desc"))
	items = append(items, NewItem("name", "desc"))
	items = append(items, NewItem("name1", "desc"))

	for i, _ := range items {
		err := items[i].Validate(reg)
		if err != nil {
			t.Error(err)
		}
	}

	items = append(items, NewItem("", "desc"))
	err := items[3].Validate(reg)
	if err == nil {
		t.Error("should fail with no name")
	}
}

func TestOutputItems(t *testing.T) {
	items := NewItems()

	items = append(items, NewItem("name", "desc"))
	items = append(items, NewItem("name", "desc"))
	items = append(items, NewItem("name1", "desc"))

	var file = "/dev/null"
	if testing.Verbose() {
		file = "/dev/stdout"
	}

	err := items.Store(file)
	if err != nil {
		t.Error(err)
	}
}
