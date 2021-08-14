package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func (p *Program) loadNoteList() error {
	noteListGOB, err := p.dbHandle.GetEntry("List")
	if err != nil {
		return fmt.Errorf("error loading NoteList: %s", err.Error())
	}
	reader := bytes.NewReader(noteListGOB)
	dec := gob.NewDecoder(reader)
	err = dec.Decode(&p.noteList)
	if err != nil {
		return fmt.Errorf("error loading NoteList: %s", err.Error())
	}
	return nil
}

func (p *Program) saveNoteList() error {
	buff := bytes.Buffer{}
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(p.noteList)
	if err != nil {
		return fmt.Errorf("error saving NoteList:\n\t%s", err.Error())
	}
	err = p.dbHandle.SaveEntry("List", buff.Bytes())
	if err != nil {
		return fmt.Errorf("error loading NoteList: %s", err.Error())
	}
	return nil
}
