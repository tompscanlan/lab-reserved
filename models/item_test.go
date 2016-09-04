package models

import (
	"testing"
	"time"

	strfmt "github.com/go-openapi/strfmt"
)

func strPtr(s string) *string { return &s }

var itemTests = []struct {
	in         Item
	valid      bool
	isReserved bool
}{
	{
		Item{Name: nil},
		false,
		false,
	},
	{
		Item{Name: strPtr("name")},
		true,
		false,
	},
	{
		Item{Name: strPtr("name"), Description: "blah"},
		true,
		false,
	},

	{
		Item{Name: strPtr("name"), Description: "reserved until now + 1 minute", ReservedUntil: strfmt.DateTime(time.Now().Add(time.Minute))},
		true,
		true,
	},

	{
		Item{Name: strPtr("name"), Description: "reserved until 1 hour from now", ReservedUntil: strfmt.DateTime(time.Now().Add(time.Hour))},
		true,
		true,
	},

	{
		Item{Name: strPtr("name"), Description: "reserved until 1 day from now", ReservedUntil: strfmt.DateTime(time.Now().Add(24 * time.Hour))},
		true,
		true,
	},

	{
		Item{Name: strPtr("name"), Description: "reserved until 1 Year from now", ReservedUntil: strfmt.DateTime(time.Now().Add(24 * time.Hour * 365))},
		true,
		true,
	},
}

func TestNewItem(t *testing.T) {
	var reg strfmt.Registry

	for i, _ := range itemTests {
		err := itemTests[i].in.Validate(reg)
		if itemTests[i].valid && err != nil {
			t.Error(err)
		}
		if !itemTests[i].valid && err == nil {
			t.Error("unexpected err for item %d", i)
		}

	}
}

func TestReserved(t *testing.T) {

	for i, _ := range itemTests {

		if itemTests[i].in.isReserved() != itemTests[i].isReserved {
			t.Errorf("%d isn't reserved, but should be", i)
		}
	}
}
