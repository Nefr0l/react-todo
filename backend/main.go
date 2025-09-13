package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"api/db"

	"api/routes"
)

func main() {
	// init db
	db.InitializeDB()

	// gin
	r := gin.Default()

	routes.TaskRoutes(r)

	r.Run(":3001")
}
