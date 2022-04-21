package main

import (
	"co/note-server/src/adapter/api"
	"co/note-server/src/adapter/persistence/ram"
)

func main() {
	var server api.NoteServer = api.NewNoteController(ram.MakeInMemoryNoteRepository())

	server.InitServer()
}
