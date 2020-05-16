package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	//TODO utilizar GraphQL

	db, err := sqlx.Connect(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/todo?sslmode=disable",
	)
	if err != nil {
		log.Fatalln(err)
	}

	//TODO utilizar Gin

	http.HandleFunc("/", Create(db))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
