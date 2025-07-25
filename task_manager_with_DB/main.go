package main

import (
	"task_manager_with_DB/controllers"
	"task_manager_with_DB/data"
	"task_manager_with_DB/router"
)

func main() {
	// MongoDB URI
	mongoURI := "mongodb://localhost:27017"
	taskService := data.NewTaskService(mongoURI, "task_manager_with_DB", "tasks")

	taskController := controllers.NewTaskController(taskService)
	r := router.SetupRouter(taskController)

	r.Run(":8080")
}
