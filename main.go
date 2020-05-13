package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sandysnunes/todo-server/src/controllers"
	"log"
	"net/http"
)

func main() {

	db, err := sqlx.Connect(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/todo?sslmode=disable",
	)
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", controllers.Create(db))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
