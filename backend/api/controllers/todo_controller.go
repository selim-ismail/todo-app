package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"todo-app/api/models"
)

func (server *Server) CreateTodoList(c *gin.Context) {
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	todoList := models.TodoList{}
	err = json.Unmarshal(body, &todoList)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	//uid, err := auth.ExtractTokenID(c.Request)
	//if err != nil {
	//	errList["Unauthorized"] = "Unauthorized"
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"status": http.StatusUnauthorized,
	//		"error":  errList,
	//	})
	//	return
	//}
	//
	//// check if the user exist:
	//user := models.User{}
	//err = server.Client.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	//if err != nil {
	//	errList["Unauthorized"] = "Unauthorized"
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"status": http.StatusUnauthorized,
	//		"error":  errList,
	//	})
	//	return
	//}
	//
	//post.AuthorID = uid //the authenticated user is the one creating the post

	todoList.Prepare()
	errorMessages := todoList.Validate()
	if len(errorMessages) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	todoListCreated, err := todoList.CreateTodoList(server.Context, server.Database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": todoListCreated,
	})
}

func (server *Server) GetAllTodoLists(c *gin.Context) {

	todoList := models.TodoList{}

	todoLists, err := todoList.FindAllTodoLists(server.Context, server.Database)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errors.New("unknown error").Error(),
		})
		return
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
