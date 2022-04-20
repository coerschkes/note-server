package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorHandler func(c *gin.Context) error

func (fn errorHandler) handleHttp(c *gin.Context) {
	if err := fn(c); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error occurred"})
	}
}
