package api

import (
	"co/note-server/src/main/config"
	"co/note-server/src/main/domain/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	repository model.NoteRepository
	properties config.ConfigProvider
}

func NewNoteController(repository model.NoteRepository) NoteController {
	return NoteController{repository: repository, properties: config.MakeConfigProvider()}
}

func (h NoteController) InitServer() {
	host := h.properties.GetProperty("server.host")
	port := h.properties.GetProperty("server.port")
	trustedProxies := h.properties.GetProperty("server.trusted.proxies")

	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies(strings.Split(trustedProxies, ","))
	router.GET("/notes", errorHandler(h.GetNotes).handleRequest)
	router.GET("/notes/:id", errorHandler(h.GetNoteById).handleRequest)
	router.POST("/notes", errorHandler(h.PostNote).handleRequest)
	router.DELETE("/notes/:id", errorHandler(h.DeleteNote).handleRequest)

	router.Run(host + ":" + port)
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
		return NewServerError(err, c.FullPath(), http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusCreated, newNote)
		return nil
	}
}

func (h NoteController) DeleteNote(c *gin.Context) *ServerError {
	paramId := c.Param("id")
	if err := h.repository.DeleteById(paramId); err != nil {
		return NewServerError(err, c.FullPath(), http.StatusBadRequest)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Note with id '" + paramId + "' deleted."})
	return nil
}
