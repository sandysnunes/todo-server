package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
)

// TodoController define the methods to perform operations on the to-do itens
type TodoController struct {
}

// NewTodoController create a instance of TodoController
func NewTodoController() *TodoController {
	return &TodoController{}
}

// FindByID get a to-do item by id
func (todoController *TodoController) FindByID(db *sqlx.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		todo := new(Todo)
		err := db.Get(todo, "SELECT id, title, description, favorite, completed FROM todo where id = $1", id)

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

// FindAll is the method to get all the to-dos from database
func (todoController *TodoController) FindAll(db *sqlx.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, description, favorite, completed FROM todo")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		var todos []Todo

		for rows.Next() {
			var id int
			var title string
			var description string
			var favorite bool
			var completed bool

			err = rows.Scan(&id, &title, &description, &favorite, &completed)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			}

			todo := Todo{ID: id, Title: title, Completed: completed, Description: description, Favorite: favorite}
			todos = append(todos, todo)
		}

		c.JSON(http.StatusOK, gin.H{"data": todos})
	}
}

// Create create a new to do item
func (todoController *TodoController) Create(db *sqlx.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		todo := new(Todo)
		err := c.Bind(todo)
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
