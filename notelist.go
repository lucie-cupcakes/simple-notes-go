package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"strings"
)

// NoteList is a dictionary for keeping track of the total Note instances
type NoteList struct {
	Value map[string]string
}

// Initialize sets up the initial values for the NoteList instance
func (l *NoteList) Initialize() {
	l.Value = make(map[string]string)
}

// Get allows to recover a Value using a Key
func (l *NoteList) Get(key string) (string, error) {
	val, found := l.Value[key]
	if found {
		return val, nil
	}
	return "", errors.New("cannot get KeyValuePair: not found")
}

// Put allows to insert a KeyValuePair
func (l *NoteList) Put(key string, value string) {
	l.Value[key] = value
}

// Has returns a bool to represent if the Key exists
func (l *NoteList) Has(key string) bool {
	_, found := l.Value[key]
	return found
}

// Count returns the amount of KeyValuePairs
func (l *NoteList) Count() int {
	return len(l.Value)
}

// ToString converts the NoteList to display it to the user
func (l *NoteList) ToString(format string) string {
	res := strings.Builder{}
	for key, value := range l.Value {
		entry := strings.ReplaceAll(format, "{key}", key)
		entry = strings.ReplaceAll(entry, "{value}", value)
		res.WriteString(entry + "\n")
	}
	return strings.TrimSpace(res.String())
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
