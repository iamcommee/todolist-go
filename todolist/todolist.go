package todolist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type database interface {
	Create(Todolist) (Todolist, error)
	GetAll() (Todolists, error)
	Update(string, Todolist) (Todolist, error)
	Delete(string) (bool, error)
}

type TodolistHandler struct {
	database database
}

func NewTodolistHandler(d database) *TodolistHandler {
	return &TodolistHandler{
		database: d,
	}
}

func (t *TodolistHandler) Create(c *gin.Context) {
	todolist := &Todolist{}
	c.ShouldBind(todolist)

	result, err := t.database.Create(*todolist)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unsuccessful created",
			"error":   "FAILED_TO_CREATE",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successful created",
		"data":    result,
	})
}

func (t *TodolistHandler) GetAll(c *gin.Context) {
	todolist, err := t.database.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unsuccessful get todolists",
			"error":   "FAILED_TO_GET",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    todolist.Todolists,
	})
}

func (t *TodolistHandler) Update(c *gin.Context) {
	todolist := &Todolist{}
	c.ShouldBindJSON(todolist)

	result, err := t.database.Update(c.Param("id"), *todolist)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unsuccessful updated",
			"error":   "ID_NOT_FOUND",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful updated",
		"data":    result,
	})
}

func (t *TodolistHandler) Delete(c *gin.Context) {
	todolist := &Todolist{}
	c.ShouldBindJSON(todolist)

	result, err := t.database.Delete(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unsuccessful deleted",
			"error":   "ID_NOT_FOUND",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful deleted",
		"data":    result,
	})
}
