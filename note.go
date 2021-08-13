package main

import (
	"bytes"
	"encoding/gob"
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
		return nil, err
	}
	return buff.Bytes(), nil
}

// FromGOB is used to deserialize a GOB byte array to a Note instance
func (n *Note) FromGOB(data []byte) error {
	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(n)
	if err != nil {
		return err
	}
	return nil
}
