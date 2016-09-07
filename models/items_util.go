package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func NewItems() Items {
	items := make(Items)
	return items
}

// returns JSON string
func (items Items) String() string {
	b, err := json.Marshal(items)

	if err != nil {
		log.Println(err)
	}

	n := bytes.IndexByte(b, 0)
	s := string(b[:n])
	return s
}

// Load items from JSON.
// clobbers current Items
func (items *Items) LoadJSON(s string) error {
	err := json.Unmarshal([]byte(s), items)
	if err != nil {
		return err
	}

	return nil
}

// Store all the items as JSON in a file
func (items Items) Store(filename string) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
