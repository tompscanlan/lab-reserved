package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func NewItems() Items {
	var items Items
	return items
}

func (items Items) String() string {
	b, err := json.Marshal(items)

	if err != nil {
		log.Println(err)
	}

	n := bytes.IndexByte(b, 0)
	s := string(b[:n])
	return s
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
