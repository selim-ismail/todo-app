package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) GetTodoById(c *gin.Context) {
	todoId := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": todoId,
	})

}
