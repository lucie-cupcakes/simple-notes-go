package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Note represents a note entry
type Note struct {
	ID           uuid.UUID
	Title        string
	Contents     string
	LastModified time.Time
	CreationTime time.Time
}

// Create sets the initial status of a Note instance
func (n *Note) Create(title string, contents string) {
	n.ID = uuid.New()
	n.CreationTime = time.Now().UTC()
	n.LastModified = n.CreationTime
	n.Title = title
	n.Contents = contents
}

// Modify is used to change the contents of a Note instance
// and have the change registered
func (n *Note) Modify(title string, contents string) {
	n.Title = title
	n.Contents = contents
	n.LastModified = time.Now().UTC()
}

// ToGOB is used to serialize the Note instance as a GOB byte array
func (n *Note) ToGOB() ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(n)
	if err != nil {
		return nil, fmt.Errorf("error serializing Note to GOB:\n\t%s", err.Error())
	}
	return buff.Bytes(), nil
}

// FromGOB is used to deserialize a GOB byte array to a Note instance
func (n *Note) FromGOB(data []byte) error {
	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		return fmt.Errorf("error restoring Note from GOB data:\n\t%s", err.Error())
	}
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(n)
	if err != nil {
		return fmt.Errorf("error restoring Note from GOB data:\n\t%s", err.Error())
	}
	return nil
}

// Save allows the user to save the Note into a Pepino Database
func (n *Note) Save(db *PepinoDB) error {
	gobBytes, err := n.ToGOB()
	if err != nil {
		return fmt.Errorf("error saving Note:\n\t%s", err.Error())
	}
	err = db.SaveEntry(n.ID.String(), gobBytes)
	if err != nil {
		return fmt.Errorf("error saving Note:\n\t%s", err.Error())
	}
	return nil
}

// Load allows the user to load the Note from a Pepino Database
func (n *Note) Load(id string, db *PepinoDB) error {
	data, err := db.GetEntry(id)
	if err != nil {
		return fmt.Errorf("cannot load Note:\n\t%s", err.Error())
	}
	err = n.FromGOB(data)
	if err != nil {
		return fmt.Errorf("cannot load Note:\n\t%s", err.Error())
	}
	return nil
}

// ToString converts the Note to a String, intended to be displayed to the user
func (n *Note) ToString() string {
	res := strings.Builder{}
	res.WriteString("Title: " + n.Title + "\n")
	res.WriteString("Creation Time: " + n.CreationTime.String() + "\n")
	res.WriteString("Last Modified: " + n.LastModified.String() + "\n")
	res.WriteString(n.Contents)
	return res.String()
}
