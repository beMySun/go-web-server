package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler handles the ping request
func (r *Registry) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
