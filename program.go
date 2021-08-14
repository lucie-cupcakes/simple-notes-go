package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Program represents the current program and its state
type Program struct {
	dbHandle  *PepinoDB
	noteList  *NoteList
	cmdInput  string
	cmdReader *bufio.Reader
}

// Initialize sets up the initial values for the Program instance
func (p *Program) Initialize() error {
	db := PepinoDB{}
	db.Initialize("http://localhost:50200", "Notes.go", "caipiroska")
	p.dbHandle = &db
	nl := NoteList{}
	nl.Initialize()
	nl.Load(p.dbHandle)
	p.noteList = &nl
	p.cmdReader = bufio.NewReader(os.Stdin)
	return nil
}

func (p *Program) readUntilFinish() string {
	fmt.Println("TIP: type @finish@ when you end the Note.")
	contentsSb := strings.Builder{}
	for {
		line, _ := p.cmdReader.ReadString('\n')
		if strings.HasPrefix(line, "@finish@") {
			break
		}
		contentsSb.WriteString(line)
	}
	return contentsSb.String()
}

func (p *Program) newCommand() {
	fmt.Print("title: ")
	title, _ := p.cmdReader.ReadString('\n')
	contents := p.readUntilFinish()
	note := Note{}
	note.Create(title, contents)
	err := note.Save(p.dbHandle)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	p.noteList.Value[note.ID.String()] = note.Title
	err = p.noteList.Save(p.dbHandle)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (p *Program) deleteCommand() {

}

func (p *Program) modifyCommand() {

}

func (p *Program) printCommand() {

}

func (p *Program) listCommand() {
	if len(p.noteList.Value) > 0 {
		for noteID, noteTitle := range p.noteList.Value {
			fmt.Println(noteID + "\t" + noteTitle)
		}
	} else {
		fmt.Println("There are not saved notes.")
	}
}

func (p *Program) helpCommand() {
	fmt.Println("Commands:\n" +
		"new-\tCreate a Note\n" +
		"del-\tDelete a Note\n" +
		"mod-\tModify a Note\n" +
		"list-\tList Notes\n" +
		"print-\tPrint a Note to the screen\n" +
		"exit-\tLeave the program.")
}

// Run is the main function for the Program object
func (p *Program) Run() {
	fmt.Println("Welcome to the Notes program!\n" +
		"TIP: type help for the command list.")
	for {
		fmt.Print("Notes>")
		p.cmdInput, _ = p.cmdReader.ReadString('\n')
		if strings.HasPrefix(p.cmdInput, "new") {
			p.newCommand()
		} else if strings.HasPrefix(p.cmdInput, "del") {
			p.deleteCommand()
		} else if strings.HasPrefix(p.cmdInput, "mod") {
			p.modifyCommand()
		} else if strings.HasPrefix(p.cmdInput, "list") {
			p.listCommand()
		} else if strings.HasPrefix(p.cmdInput, "print") {
			p.printCommand()
		} else if strings.HasPrefix(p.cmdInput, "help") {
			p.helpCommand()
		} else if strings.HasPrefix(p.cmdInput, "exit") {
			break
		}
	}
}

func main() {
	program := Program{}
	err := program.Initialize()
	if err != nil {
		panic(err)
	}
	program.Run()
}
