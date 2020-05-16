package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

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
