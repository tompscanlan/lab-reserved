package models

import (
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
)

var resTests = []struct {
	in      Reservation
	valid   bool
	other   Reservation
	overlap bool
}{
	{
		NewReservation("tom", time.Now(), time.Now().Add(3*time.Hour)),
		true,
		NewReservation("bob", time.Now(), time.Now().Add(3*time.Hour)),
		true,
	},
	{
		NewReservation("tom", time.Now(), time.Now().Add(3*time.Hour)),
		true,
		NewReservation("bob", time.Now().Add(4*time.Hour), time.Now().Add(6*time.Hour)),
		false,
	},
}

func TestReservationValid(t *testing.T) {
	var reg strfmt.Registry

	for i, test := range resTests {
		err := test.in.Validate(reg)
		if test.valid == (err != nil) {
			t.Errorf("%d validate should be (%v), but isn't: %s", i, test.valid, err)
		}
	}
}

func TestReservationOverlap(t *testing.T) {
	for i := range resTests {
		test := resTests[i]

		err, overlap := test.in.Overlap(test.other)
		if err != nil {
			t.Error(err)
			continue
		}
		if overlap != test.overlap {
			t.Errorf("%d overlap should be (%v) but is (%v)", i, test.overlap, overlap)
		}
	}
}
