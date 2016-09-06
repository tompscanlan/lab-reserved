package models

import (
	"log"
	"testing"
	"time"

	strfmt "github.com/go-openapi/strfmt"
)

func strPtr(s string) *string                          { return &s }
func resPtr(s string, t time.Time, i int) *Reservation { n := NewReservation(s, t, i); return &n }

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
		Item{Description: "reserved at same time", ID: 0,
			Name: strPtr("server4"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), 1),
				resPtr("bob", time.Now(), 1)}},
		true,
		true,
	},

	{
		Item{Description: "same time, different lengths", ID: 0,
			Name: strPtr("server5"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), 3),
				resPtr("bob", time.Now(), 12)},
		},
		true,
		true,
	},
	{
		Item{Description: "several same time, different lengths", ID: 0,
			Name: strPtr("server6"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), 3),
				resPtr("bob", time.Now(), 12),
				resPtr("tim", time.Now().Add(3*time.Hour), 12),
				resPtr("babs", time.Now().Add(7*time.Minute), 12)},
		},
		true,
		true,
	},

	{
		Item{Description: "", ID: 0,
			Name: strPtr("server10"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), 1),
				resPtr("bill", time.Now(), 1)}},
		true,
		true,
	},

	{
		Item{Description: "non-overlap", ID: 0,
			Name: strPtr("server11"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now().Add(1*time.Hour), 1),
				resPtr("bill", time.Now().Add(2*time.Hour), 1)}},
		true,
		false,
	},
	{
		Item{Description: "several non-overlap", ID: 0,
			Name: strPtr("server11"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), 2),
				resPtr("tad", time.Now().Add(3*time.Hour), 3),
				resPtr("bill", time.Now().Add(7*time.Hour), 1)}},
		true,
		true,
	},
}

func TestNewItem(t *testing.T) {
	var reg strfmt.Registry

	for i, test := range itemTests {
		test.in.Reserve("tom", time.Now(), 3)
		test.in.Reserve("tom", time.Now().Add(4), 3)
		err := test.in.Validate(reg)
		if test.valid && err != nil {
			t.Error(err)
		}
		if !test.valid && err == nil {
			t.Error("unexpected err for item %d", i)
		}
		log.Println(test.in)

		log.Printf("%#v", test.in)

	}

}

func TestReserved(t *testing.T) {
	time.Sleep(2 * time.Second)
	for i, _ := range itemTests {
		if itemTests[i].in.isReserved() != itemTests[i].isReserved {
			t.Errorf("%d isn't reserved, but should be", i)
		}
	}
}
