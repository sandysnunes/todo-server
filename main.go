package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()

	todoController := TodoController{}

	ginEngine.GET("/todo", todoController.FindAll())
	ginEngine.GET("/todo/:id", todoController.FindById())

	ginEngine.Run(":8080")

}
