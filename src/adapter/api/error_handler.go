package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type errorHandler func(c *gin.Context) *ServerError

func (fn errorHandler) handleRequest(c *gin.Context) {
	if err := fn(c); err != nil {
		fmt.Println("[ERROR] on path: '"+err.Path+"' |", err.Error)
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
	}
}
