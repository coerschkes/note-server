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
	router := gin.Default()
	router.GET("/notes", h.GetNotes)
	router.GET("/notes/:id", h.GetNoteById)
	router.POST("/notes", h.PostNote)
	router.DELETE("/notes/:id", h.DeleteNote)

	router.Run("localhost:8080")
}

func (h HttpHandler) GetNotes(c *gin.Context) {
	note, err := h.repository.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error occurred"})
	} else {
		c.IndentedJSON(http.StatusOK, note)
	}
}

func (h HttpHandler) GetNoteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	note, err := h.repository.GetById(int64(id))

	if err == nil {
		c.IndentedJSON(http.StatusOK, note)
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing param 'id'"})
	}
}

func (h HttpHandler) PostNote(c *gin.Context) {
	var newNote note.Note

	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	h.repository.Add(newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func (h HttpHandler) DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	err = h.repository.DeleteById(int64(id))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing param 'id'"})
	}
}
