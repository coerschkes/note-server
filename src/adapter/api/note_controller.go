package api

import (
	"co/note-server/src/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const trustedProxies = "0.0.0.0"

type NoteController struct {
	repository model.NoteRepository
}

func NewNoteController(repository model.NoteRepository) NoteController {
	return NoteController{repository: repository}
}

func (h NoteController) InitServer() {
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies([]string{trustedProxies})
	router.GET("/notes", errorHandler(h.GetNotes).handleRequest)
	router.GET("/notes/:id", errorHandler(h.GetNoteById).handleRequest)
	router.POST("/notes", errorHandler(h.PostNote).handleRequest)
	router.DELETE("/notes/:id", errorHandler(h.DeleteNote).handleRequest)

	router.Run("0.0.0.0:8080")
}

func (h NoteController) GetNotes(c *gin.Context) *ServerError {
	if note, err := h.repository.GetAll(); err != nil {
		return NewServerError(err, c.FullPath(), http.StatusInternalServerError)
	} else {
		c.IndentedJSON(http.StatusOK, note)
		return nil
	}
}

func (h NoteController) GetNoteById(c *gin.Context) *ServerError {
	paramId := c.Param("id")
	if note, err := h.repository.GetById(paramId); err != nil {
		return NewServerError(err, c.FullPath(), http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, note)
		return nil
	}
}

func (h NoteController) PostNote(c *gin.Context) *ServerError {
	var newNote model.Note

	if err := c.BindJSON(&newNote); err != nil {
		return NewServerError(err, c.FullPath(), http.StatusBadRequest)
	}

	if err := h.repository.Add(newNote); err != nil {
		return NewServerError(err, c.FullPath(), http.StatusInternalServerError)
	} else {
		c.IndentedJSON(http.StatusCreated, newNote)
		return nil
	}
}

func (h NoteController) DeleteNote(c *gin.Context) *ServerError {
	paramId := c.Param("id")
	if err := h.repository.DeleteById(paramId); err != nil {
		return NewServerError(err, c.FullPath(), http.StatusInternalServerError)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Note with id '" + paramId + "' deleted."})
	return nil
}
