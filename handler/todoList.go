package handler

import (
	"go-web-server/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

// QueryTodoList return ListTodoList
func (r *Registry) QueryTodoList(c *gin.Context) {
	result := r.DAO.QueryTodoList()
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"code":    1,
		"message": "Query Success",
	})
}

// DeleteTodo will delete a record
func (r *Registry) DeleteTodo(c *gin.Context) {
	var data dao.TodoList

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data structure mismatched: " + err.Error(),
		})
		return
	}

	if err := r.DAO.DeleteTodo(data.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
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

// UpdateteTodoRecord create a new record
func (r *Registry) UpdateteTodoRecord(c *gin.Context) {
	var data dao.TodoList
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data structure mismatched: " + err.Error(),
		})
		return
	}

	if err := r.DAO.UpdateTodo(data.ID, data.IsCompleted); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Report Success",
	})

}
