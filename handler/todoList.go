package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListTodoList return ListTodoList
func (r *Registry) ListTodoList(c *gin.Context) {
	result := r.DAO.ListTodoList()
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"code":    1,
		"message": "Query Success",
	})
}
