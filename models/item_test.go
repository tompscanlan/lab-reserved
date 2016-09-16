package models

import (
	"log"
	"testing"
	"time"

	strfmt "github.com/go-openapi/strfmt"
)

func strPtr(s string) *string                                { return &s }
func resPtr(s string, t time.Time, e time.Time) *Reservation { n := NewReservation(s, t, e); return &n }

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
				resPtr("tom", time.Now(), time.Now().Add(6*time.Hour)),
				resPtr("bob", time.Now(), time.Now().Add(12*time.Hour))}},
		true,
		true,
	},

	{
		Item{Description: "same time, different lengths", ID: 0,
			Name: strPtr("server5"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), time.Now().Add(3*time.Hour)),
				resPtr("bob", time.Now(), time.Now().Add(12*time.Hour))},
		},
		true,
		true,
	},
	{
		Item{Description: "several same time, different lengths", ID: 0,
			Name: strPtr("server6"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), time.Now().Add(3*time.Hour)),
				resPtr("bob", time.Now(), time.Now().Add(12*time.Hour)),
				resPtr("tim", time.Now().Add(3*time.Hour), time.Now().Add(12*time.Hour)),
				resPtr("babs", time.Now().Add(7*time.Minute), time.Now().Add(12*time.Hour))},
		},
		true,
		true,
	},

	{
		Item{Description: "", ID: 0,
			Name: strPtr("server10"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), time.Now().Add(1*time.Hour)),
				resPtr("bill", time.Now(), time.Now().Add(1*time.Hour))}},
		true,
		true,
	},

	{
		Item{Description: "non-overlap", ID: 0,
			Name: strPtr("server11"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now().Add(1*time.Hour), time.Now().Add(1*time.Hour)),
				resPtr("bill", time.Now().Add(2*time.Hour), time.Now().Add(1*time.Hour))}},
		true,
		false,
	},
	{
		Item{Description: "several non-overlap", ID: 0,
			Name: strPtr("server11"),
			Reservations: []*Reservation{
				resPtr("tom", time.Now(), time.Now().Add(2*time.Hour)),
				resPtr("tad", time.Now().Add(3*time.Hour), time.Now().Add(3*time.Hour)),
				resPtr("bill", time.Now().Add(7*time.Hour), time.Now().Add(1*time.Hour))}},
		true,
		true,
	},
}

func TestNewItem(t *testing.T) {
	var reg strfmt.Registry

	for i, test := range itemTests {
		test.in.Reserve("tom", time.Now(), time.Now().Add(3*time.Hour))
		test.in.Reserve("tom", time.Now().Add(4), time.Now().Add(3*time.Hour))
		err := test.in.Validate(reg)
		if test.valid && err != nil {
			t.Error(err)
		}
		if !test.valid && err == nil {
			t.Errorf("unexpected err for item %d", i)
		}
		log.Println(test.in)
		//		log.Printf("%#v", test.in)

	}

}

func TestReserved(t *testing.T) {
	time.Sleep(2 * time.Second)
	for i := range itemTests {
		if itemTests[i].in.isReserved() != itemTests[i].isReserved {
			t.Errorf("%d isn't reserved, but should be", i)
		}
	}
}
