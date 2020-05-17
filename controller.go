package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
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
			if errors.Is(err, sql.ErrNoRows) {
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

func (todoController *TodoController) Create(db *sqlx.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var todo Todo
		err := c.Bind(&todo)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		stmt, err := db.Prepare("INSERT INTO todo(title, description, favorite) VALUES($1, $2, $3) returning id")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		defer stmt.Close()

		if _, err := stmt.Exec(todo.Title, todo.Description, todo.Favorite); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		} else {
			c.JSON(http.StatusCreated, todo)
		}

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
