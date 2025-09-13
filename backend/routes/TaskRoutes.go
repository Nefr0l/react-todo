package routes

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	tasks := r.Group("/tasks")
	{
		tasks.GET("/", handlers.GetTasks)
		tasks.GET("/:id", handlers.GetTask)
		tasks.POST("/", handlers.CreateTask)
		tasks.DELETE("/:id", handlers.DeleteTask)

	}
}
