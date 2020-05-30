// TODO service and repository ?
// TODO Dependecy injection ?
// TODO gorm orm ?

package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	// HOST used by Gin
	HOST = "localhost"
	
	// PORT used by Gin
    PORT = "7777"
)

func main() {

	dataSource := "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable"
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatalln(err)
	}

	route := gin.Default()
	route.Use(cors.Default())

	todoController := NewTodoController()

	route.GET("/todo", todoController.FindAll(db))
	route.GET("/todo/:id", todoController.FindByID(db))
	route.POST("/todo", todoController.Create(db))

	err = route.Run(HOST + ":" + PORT)
	log.Println("Running @ http://" + HOST + ":" + PORT )
	if err != nil {
		log.Fatalln(err)
	}
}
