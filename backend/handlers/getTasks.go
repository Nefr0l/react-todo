package handlers

import (
	"api/db"
	"api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	rows, err := db.DB.Query("select * from tasks;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.Id, &task.Name, &task.IsCompleted)
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	row := db.DB.QueryRow("select * from tasks where id = ?", id)

	var task models.Task
	err := row.Scan(&task.Id, &task.Name, &task.IsCompleted)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, task)
}
