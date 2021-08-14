package main

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

func main() {
	program := Program{}
	err := program.Initialize()
	if err != nil {
		panic(err)
	}
}
