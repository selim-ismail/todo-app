package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/api/models"
)

func (server *Server) GetAllTodoLists(c *gin.Context) {

	todoList := models.TodoList{}

	todoLists, err := todoList.FindAllTodoLists()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errors.New("unknown error").Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": todoLists,
	})
}

func (server *Server) GetTodoById(c *gin.Context) {
	todoId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": todoId,
	})

}
