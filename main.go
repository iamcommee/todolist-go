package main

import (
	"todolist/database"
	"todolist/todolist"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database := database.NewFileDatabase()
	// database := database.NewMongoDatabase()

	handler := todolist.NewTodolistHandler(database)

	r.POST("/todolists", handler.Create)
	r.GET("/todolists", handler.GetAll)
	r.PATCH("/todolists/:id", handler.Update)
	r.DELETE("/todolists/:id", handler.Delete)

	r.Run(":8888")
}
