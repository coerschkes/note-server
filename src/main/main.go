package main

import (
	"co/note-server/src/main/adapter/api"
	"co/note-server/src/main/adapter/persistence/db"
)

func main() {
	var server api.NoteServer = api.NewNoteController(db.MakeRedisRepository())

	server.InitServer()
}
