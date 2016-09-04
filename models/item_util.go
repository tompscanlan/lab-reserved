package models

import (
	"log"
	"time"

	"github.com/go-openapi/strfmt"
)

//func NewItem() *Item {
//	item := new(Item)
//	return item
//}

func NewItem(name string, description string) *Item {
	item := new(Item)
	item.Name = &name
	item.Description = description
	return item
}

func (item *Item) Reserve(owner string, date strfmt.Date) {

}

//strfmt.DateTime(time.Now().Add(time.Hour))},

func (item Item) isReservedOn(date time.Time) bool {

	parsed, err := time.Parse(time.RFC3339, item.ReservedUntil.String())
	if err != nil {
		log.Panicln(err)
		return false
	}
	//	log.Printf("checking reservation for date: %s, and item is reserved until %s", parsed, date)

	//log.Println("difference in dates: ", date.Sub(parsed))
	// if reserved is in the future
	if date.Sub(parsed) < 0 {
		return true
	}
	return false
}

func (item Item) isReserved() bool {
	date := time.Now()
	return item.isReservedOn(date)

}
