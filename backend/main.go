package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {

	// init db
	var err error
	db, err = sql.Open("sqlite3", "./db/tasks.db")

	if err != nil {
		log.Fatal("error loading database")
	}

	// gin
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("tasks", getTasks)
	}

	r.Run(":3001")
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getTasks Called"})
}
