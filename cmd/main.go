package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "admin"
	cfg.Net = "tcp"
	cfg.Addr = "localhost:3307"
	cfg.DBName = "database"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	handleErr(err)
	fmt.Println("connected to the database")

	// test query
	rows, err := db.Query("select id, name from tasks;")
	handleErr(err)

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)

		handleErr(err)
		fmt.Println(id, name)
	}

}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
