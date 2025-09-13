package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("sqlite3", "./db/app.sqlite")
	handleErr(err)

	defer db.Close()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rows, _ := db.Query("SELECT id, name FROM tasks")
			defer rows.Close()
		}
	})

	log.Println("Server running on :3001")
	http.ListenAndServe(":3001", nil)

}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
