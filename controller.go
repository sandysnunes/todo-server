package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type TodoController struct {
}

func (todoController *TodoController) FindById(db *sqlx.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var todo = Todo{}
		err := db.Get(&todo, "SELECT id, title, description, favorite FROM todo where id = $1", id)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"data": []string{}})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"data": todo})
		}
	}
}

func (todoController *TodoController) FindAll(db *sqlx.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, description, favorite FROM todo")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		var todos []Todo

		for rows.Next() {
			var id int
			var title string
			var description sql.NullString
			var favorite bool
			err = rows.Scan(&id, &title, &description, &favorite)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			}
			todos = append(todos, Todo{id, title, title, favorite, nil})
		}

		c.JSON(http.StatusOK, gin.H{"data": todos})
	}
}

/*func Create(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := db.QueryRow("INSERT INTO todo(title, description) VALUES ($1, $2)", "titulo", "descricao")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal error"))
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}*/
