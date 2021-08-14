package main

// Program represents the current program and its state
type Program struct {
	dbHandle     *PepinoDB
	noteList     map[string]string
	commandInput string
}

// Initialize sets up the initial values for the Program instance
func (p *Program) Initialize() error {
	var db PepinoDB
	db.Initialize("http://localhost:50200", "Notes.go", "caipiroska")
	p.dbHandle = &db
	if p.loadNoteList() != nil {
		p.noteList = make(map[string]string)
	}
	return nil
}

func main() {
	program := Program{}
	err := program.Initialize()
	if err != nil {
		panic(err)
	}
}
