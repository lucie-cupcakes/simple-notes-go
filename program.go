package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Program represents the current program and its state
type Program struct {
	dbHandle     *PepinoDB
	noteList     *NoteList
	commandInput string
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
	return nil
}

func (p *Program) newCommand() {

}

func (p *Program) deleteCommand() {

}

func (p *Program) modifyCommand() {

}

func (p *Program) printCommand() {

}

func (p *Program) listCommand() {

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

func (p *Program) readUntilFinish(reader *bufio.Reader) {
	contentsSb := strings.Builder{}
	for {
		line, _ := reader.ReadString('\n')
		if strings.HasPrefix(line, "@finish@") {
			break
		}
		contentsSb.WriteString(line)
	}
}

// Run is the main function for the Program object
func (p *Program) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Notes program!\n" +
		"TIP: type help for the command list.")
	for {
		p.commandInput, _ = reader.ReadString('\n')
		if strings.HasPrefix(p.commandInput, "new") {
			p.newCommand()
		} else if strings.HasPrefix(p.commandInput, "del") {
			p.deleteCommand()
		} else if strings.HasPrefix(p.commandInput, "mod") {
			p.modifyCommand()
		} else if strings.HasPrefix(p.commandInput, "list") {
			p.listCommand()
		} else if strings.HasPrefix(p.commandInput, "print") {
			p.printCommand()
		} else if strings.HasPrefix(p.commandInput, "help") {
			p.helpCommand()
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
