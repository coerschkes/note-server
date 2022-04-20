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

func (h HttpHandler) GetNotes(c *gin.Context) error {
	if note, err := h.repository.GetAll(); err != nil {
		return err
	} else {
		c.IndentedJSON(http.StatusOK, note)
		return nil
	}
}

func (h HttpHandler) GetNoteById(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if note, err := h.repository.GetById(int64(id)); err != nil {
		return err
	} else {
		c.IndentedJSON(http.StatusOK, note)
		return nil
	}
}

func (h HttpHandler) PostNote(c *gin.Context) error {
	var newNote note.Note

	if err := c.BindJSON(&newNote); err != nil {
		return err
	}

	if err := h.repository.Add(newNote); err != nil {
		return err
	} else {
		c.IndentedJSON(http.StatusCreated, newNote)
		return nil
	}
}

func (h HttpHandler) DeleteNote(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	return h.repository.DeleteById(int64(id))
}
