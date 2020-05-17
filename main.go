package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	dataSource := "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable"

	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatalln(err)
	}

	ginEngine := gin.Default()
	todoController := &TodoController{}

	ginEngine.GET("/todo", todoController.FindAll(db))
	ginEngine.GET("/todo/:id", todoController.FindById(db))
	ginEngine.POST("/todo", todoController.Create(db))

	err = ginEngine.Run()

	if err != nil {
		log.Fatalln(err)
	}
}
