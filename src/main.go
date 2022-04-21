package main

import (
	"co/note-server/src/adapter/api"
	"co/note-server/src/adapter/persistence"
)

func main() {
	var server api.NoteServer = api.NewNoteController(persistence.MakeInMemoryNoteRepository())

	server.InitServer()
}
