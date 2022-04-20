package main

import (
	"co/note-server/src/network"
	"co/note-server/src/note"
)

func main() {
	var server network.NoteServer = network.NewHttpHandler(note.MakeInMemoryNoteRepository())

	server.Init()
}
