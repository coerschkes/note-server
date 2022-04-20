package network

import "github.com/gin-gonic/gin"

type NoteServer interface {
	GetNotes(c *gin.Context)
	GetNoteById(c *gin.Context)
	PostNote(c *gin.Context)
	DeleteNote(c *gin.Context)
	Init()
}
