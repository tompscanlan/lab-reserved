package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func NewUsers() Users {
	users := make(Users)
	return users
}

func (users Users) String() string {
	b, err := json.Marshal(users)

	if err != nil {
		log.Println(err)
	}

	n := bytes.IndexByte(b, 0)
	s := string(b[:n])
	return s
}

// Store all the users as JSON in a file
func (users Users) Store(filename string) error {
	b, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
