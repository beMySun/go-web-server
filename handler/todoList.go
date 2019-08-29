package handler

import (
	"go-web-server/dao"
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

// CreateTodoRecord create a new record
func (r *Registry) CreateTodoRecord(c *gin.Context) {
	var data dao.TodoList
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data structure mismatched: " + err.Error(),
		})
		return
	}

	if err := r.DAO.CreateTodoRecord(data.ID, data.Text, data.IsCompleted); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Report Success",
	})

}
