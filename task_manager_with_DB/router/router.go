package router

import (
	"task_manager_with_DB/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", taskController.CreateTask)
	r.GET("/tasks", taskController.GetAllTasks)
	r.GET("/tasks/:id", taskController.GetTaskByID)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}
