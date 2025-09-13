package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var err error
var rows *sql.Rows

func InitializeDB() {
	DB, err = sql.Open("sqlite3", "db/todo.db")
	checkError()

	_, err = DB.Exec(`
		create table if not exists tasks (
			id integer primary key autoincrement,
			name varchar(255),
			isCompleted boolean
		);
	`)
	checkError()

	_, err = DB.Exec(`
		insert into tasks values (null, 'test', true);
	`)
	checkError()

	log.Print(rows)
}

func checkError() {
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("[OK]")
	}
}
