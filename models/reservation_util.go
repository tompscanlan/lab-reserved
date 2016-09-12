package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	//	"github.com/asaskevich/govalidator"
)

func NewReservation(username string, begin time.Time, hours int) Reservation {
	r := new(Reservation)

	h := int64(hours)
	b := strfmt.DateTime(begin)
	r.Username = &username
	r.Begin = &b
	r.Hoursheld = &h
	return *r
}

func StrfmtDateTimeToTime(date *strfmt.DateTime) time.Time {

	parsed, err := time.Parse(time.RFC3339, date.String())
	if err != nil {
		zero := new(time.Time)

		log.Println(err)
		return *zero
	}
	return parsed
}

func (r Reservation) BeginTime() time.Time {
	return StrfmtDateTimeToTime(r.Begin)
}

func (r Reservation) EndTime() time.Time {
	return r.BeginTime().Add(time.Duration(*r.Hoursheld) * time.Hour)
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
	parsed, err := time.Parse(time.RFC3339, r.Begin.String())
	if err != nil {
		log.Println(err)
		zero := new(time.Time)
		return err, *zero
	}
	return nil, parsed
}

func (r Reservation) GetEndTime() (error, time.Time) {
	zero := new(time.Time)

	err, a := r.GetTime()
	if err != nil {
		return err, *zero
	}

	aend := a.Add(time.Duration(*r.Hoursheld) * time.Hour)

	return nil, aend
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
	//	log.Printf("a: %s--%s, b: %s--%s", a, aend, b, bend)

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
