package api

import "github.com/gin-gonic/gin"

type NoteServer interface {
	InitServer()
	GetNotes(c *gin.Context) *ServerError
	GetNoteById(c *gin.Context) *ServerError
	PostNote(c *gin.Context) *ServerError
	DeleteNote(c *gin.Context) *ServerError
}

type ServerError struct {
	Error   error
	Message string
	Path    string
	Code    int
}

func NewServerError(err error, path string, code int) *ServerError {
	return &ServerError{Error: err, Message: err.Error(), Path: path, Code: code}
}
