all: simple-notes-go
.PHONY: clean

simple-notes-go:
	go build

clean:
	test -f simple-notes-go && rm simple-notes-go

