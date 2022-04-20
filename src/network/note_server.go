package network

import "github.com/gin-gonic/gin"

type NoteServer interface {
	Init()
	GetNotes(c *gin.Context) error
	GetNoteById(c *gin.Context) error
	PostNote(c *gin.Context) error
	DeleteNote(c *gin.Context) error
}
