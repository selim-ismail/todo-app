package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetTodoById(c *gin.Context) {
	todoId := c.Param("id")
	fmt.Println(todoId)
}
