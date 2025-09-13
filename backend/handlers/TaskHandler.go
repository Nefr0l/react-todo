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

func CreateTask(c *gin.Context) {
	// curl -X POST http://localhost:3001/tasks/ \-H "Content-Type: application/json" \-d '{"name":"zadanie 2", "isCompleted":true}'

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}

	res, err := db.DB.Exec(
		"insert into tasks values (null, ?, ?)",
		task.Name, task.IsCompleted,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}

	id, _ := res.LastInsertId()
	task.Id = int(id)
	c.JSON(http.StatusCreated, task)
}

func DeleteTask(c *gin.Context) {
	// curl -X DELETE http://localhost:3001/tasks/3

	id := c.Param("id")

	_, err := db.DB.Exec("DELETE from tasks where id like ?", id)

	if err != nil {
		log.Fatal(err)
	}
}
