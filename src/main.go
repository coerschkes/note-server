package main

import (
	"co/note-server/src/note"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var repo note.NoteRepository = note.MakeInMemoryNoteRepository()

func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.GET("/notes/:id", getNoteById)
	router.POST("/notes", postNote)
	router.DELETE("/notes/:id", deleteNote)

	router.Run("localhost:8080")
}

func getNotes(c *gin.Context) {
	note, err := repo.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error occurred"})
	} else {
		c.IndentedJSON(http.StatusOK, note)
	}
}

func getNoteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	note, err := repo.GetById(int64(id))

	if err == nil {
		c.IndentedJSON(http.StatusOK, note)
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing param 'id'"})
	}
}

func postNote(c *gin.Context) {
	var newNote note.Note

	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	repo.Add(newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func deleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	err = repo.DeleteById(int64(id))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing param 'id'"})
	}
}
