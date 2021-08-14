package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// NoteList is a dictionary for keeping track of the total Note instances
type NoteList struct {
	Value map[string]string
}

// Initialize sets up the initial values for the NoteList instance
func (l *NoteList) Initialize() {
	l.Value = make(map[string]string)
}

// Load recovers the NoteList instance from the PepinoDB
func (l *NoteList) Load(dbHandle *PepinoDB) error {
	noteListGOB, err := dbHandle.GetEntry("List")
	if err != nil {
		return fmt.Errorf("error loading NoteList: %s", err.Error())
	}
	reader := bytes.NewReader(noteListGOB)
	dec := gob.NewDecoder(reader)
	err = dec.Decode(l)
	if err != nil {
		return fmt.Errorf("error loading NoteList: %s", err.Error())
	}
	return nil
}

// Save serializes and saves the NoteList instance to the PepinoDB
func (l *NoteList) Save(dbHandle *PepinoDB) error {
	buff := bytes.Buffer{}
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(l)
	if err != nil {
		return fmt.Errorf("error saving NoteList:\n\t%s", err.Error())
	}
	err = dbHandle.SaveEntry("List", buff.Bytes())
	if err != nil {
		return fmt.Errorf("error saving NoteList: %s", err.Error())
	}
	return nil
}
