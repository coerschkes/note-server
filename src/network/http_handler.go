package network

import (
	"co/note-server/src/note"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	repository note.NoteRepository
}

func NewHttpHandler(repository note.NoteRepository) HttpHandler {
	return HttpHandler{repository: repository}
}

func (h HttpHandler) Init() {
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/notes", errorHandler(h.GetNotes).handleHttp)
	router.GET("/notes/:id", errorHandler(h.GetNoteById).handleHttp)
	router.POST("/notes", errorHandler(h.PostNote).handleHttp)
	router.DELETE("/notes/:id", errorHandler(h.DeleteNote).handleHttp)

	router.Run("localhost:8080")
}

func (h HttpHandler) GetNotes(c *gin.Context) *ServerError {
	if note, err := h.repository.GetAll(); err != nil {
		return NewServerError(err, "Unable to get notes from the repository.", c.FullPath(), http.StatusInternalServerError)
	} else {
		c.IndentedJSON(http.StatusOK, note)
		return nil
	}
}

func (h HttpHandler) GetNoteById(c *gin.Context) *ServerError {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return NewServerError(err, "Unable to parse parameter 'id' from actual parameter '"+paramId+"'.", c.FullPath(), http.StatusBadRequest)
	}

	if note, err := h.repository.GetById(int64(id)); err != nil {
		return NewServerError(err, "Unable to find note with id '"+paramId+"' from the repository.", c.FullPath(), http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, note)
		return nil
	}
}

func (h HttpHandler) PostNote(c *gin.Context) *ServerError {
	var newNote note.Note

	if err := c.BindJSON(&newNote); err != nil {
		return NewServerError(err, "Unable to create note from provided json.", c.FullPath(), http.StatusBadRequest)
	}

	if err := h.repository.Add(newNote); err != nil {
		return NewServerError(err, "Unable to save the provided note.", c.FullPath(), http.StatusInternalServerError)
	} else {
		c.IndentedJSON(http.StatusCreated, newNote)
		return nil
	}
}

func (h HttpHandler) DeleteNote(c *gin.Context) *ServerError {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return NewServerError(err, "Unable to parse parameter 'id' from actual parameter '"+paramId+"'.", c.FullPath(), http.StatusBadRequest)
	}

	if err := h.repository.DeleteById(int64(id)); err != nil {
		return NewServerError(err, "Unable to delete note with id '"+paramId+"'.", c.FullPath(), http.StatusInternalServerError)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Note with id '" + paramId + "' deleted."})
	return nil
}
