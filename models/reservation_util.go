package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	//	"github.com/asaskevich/govalidator"
)

func NewReservation(username string, begin time.Time, end time.Time) Reservation {
	r := new(Reservation)

	b := strfmt.DateTime(begin)
	e := strfmt.DateTime(end)
	f := true
	r.Username = &username
	r.Begin = &b
	r.End = &e
	r.Approved = &f
	log.Printf("made new reservation: %s", r.String())
	return *r
}

func StrfmtDateTimeToTime(date *strfmt.DateTime) time.Time {
	parsed, err := time.Parse(time.RFC3339, date.String())
	if err != nil {
		log.Println(err)
		zero := new(time.Time)
		return *zero
	}
	return parsed
}

func (r Reservation) String() string {
	b, err := json.Marshal(r)

	if err != nil {
		log.Println(err)
	}

	s := string(b[:])
	return s
}

// returns a zero time on error
func (r Reservation) GetTime() (error, time.Time) {

	parsed, err := time.Parse(time.RFC3339, (*r.Begin).String())
	if err != nil {
		log.Println(err)
		zero := new(time.Time)
		return err, *zero
	}
	return nil, parsed
}

func (r Reservation) GetEndTime() (error, time.Time) {
	parsed, err := time.Parse(time.RFC3339, (*r.End).String())
	if err != nil {
		log.Println(err)
		zero := new(time.Time)
		return err, *zero
	}
	return nil, parsed
}

func (r Reservation) ReservedAt(time time.Time) (error, bool) {
	err, a := r.GetTime()
	if err != nil {
		return err, false
	}

	err, aend := r.GetEndTime()
	if err != nil {
		return err, false
	}
	if time.After(a) && time.Before(aend) {
		return nil, true
	}
	return nil, false
}
func (r Reservation) Overlap(other Reservation) (error, bool) {
	err, a := r.GetTime()
	if err != nil {
		return err, false
	}
	err, aend := r.GetEndTime()
	if err != nil {
		return err, false
	}

	err, b := other.GetTime()
	if err != nil {
		return err, false
	}

	err, bend := other.GetEndTime()
	if err != nil {
		return err, false
	}
	log.Printf("a: %s--%s, b: %s--%s", a, aend, b, bend)

	// if they start or end at same time, there is overlap
	if a.Equal(b) || aend.Equal(bend) {
		return nil, true
	}
	// a in entierly before b
	if a.Before(b) && (aend.Before(b) || aend.Equal(b)) {
		return nil, false
	}

	// b in entierly before a
	if b.Before(a) && (bend.Before(a) || bend.Equal(a)) {
		return nil, false
	}

	return nil, false
}
