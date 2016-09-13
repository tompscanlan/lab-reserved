package models

import (
	"encoding/json"
	"log"
	"time"
)

func NewItem(name string, description string) Item {
	item := new(Item)
	item.Name = &name
	item.Description = description
	return *item
}

func (item *Item) Reserve(owner string, starttime time.Time, endtime time.Time) bool {

	r := NewReservation(owner, starttime, endtime)
	item.Reservations = append(item.Reservations, &r)

	return true
}

func (item Item) isReservedOn(time time.Time) bool {

	for _, r := range item.Reservations {
		_, taken := r.ReservedAt(time)

		if taken {
			return true
		}
	}
	return false
}

func (item Item) isReserved() bool {
	date := time.Now()
	return item.isReservedOn(date)
}

func (item Item) String() string {
	b, err := json.Marshal(item)

	if err != nil {
		log.Println(err)
	}

	s := string(b[:])
	return s
}
