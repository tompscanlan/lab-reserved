package models

import (
	"testing"

	strfmt "github.com/go-openapi/strfmt"
)

func TestNewItems(t *testing.T) {
	var reg strfmt.Registry
	items := NewItems()

	items["name"] = NewItem("name", "desc")
	items["name1"] = NewItem("name1", "desc")
	items["name2"] = NewItem("name2", "desc")

	for i := range items {
		item := items[i]
		err := item.Validate(reg)
		if err != nil {
			t.Error(err)
		}
	}

	items[""] = NewItem("", "desc")
	item := items[""]
	err := item.Validate(reg)
	if err == nil {
		t.Error("should fail with no name")
	}
}

func TestOutputItems(t *testing.T) {
	items := NewItems()

	items["name"] = NewItem("name", "desc")
	items["name1"] = NewItem("name1", "desc")
	items["name2"] = NewItem("name2", "desc")

	var file = "/dev/null"
	if testing.Verbose() {
		file = "/dev/stdout"
	}

	err := items.Store(file)
	if err != nil {
		t.Error(err)
	}
}
