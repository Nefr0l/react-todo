package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/app.sqlite")

	handleErr(err)

	defer db.Close()

}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
