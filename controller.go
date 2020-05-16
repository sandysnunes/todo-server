package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type TodoController struct {
}

func (todoController *TodoController) FindById() func(context *gin.Context) {
	return func(context *gin.Context) {
		user := context.Params.ByName("id")
		value, ok := db[user]
		if ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	}
}

func (todoController *TodoController) FindAll() func(context *gin.Context) {
	return func(context *gin.Context) {
		user := context.Params.ByName("name")
		value, ok := db[user]
		if ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	}
}

func Create(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := db.QueryRow("INSERT INTO todo(title, description) VALUES ($1, $2)", "titulo", "descricao")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal error"))
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
