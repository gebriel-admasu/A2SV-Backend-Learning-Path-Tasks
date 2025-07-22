package router

import (
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)

func SetupRouter(service *data.TaskService) *gin.Engine {
	r := gin.Default()
	tc := controllers.NewTaskController(service)

	r.GET("/tasks", tc.GetAllTasks)
	r.GET("/tasks/:id", tc.GetTask)
	r.POST("/tasks", tc.CreateTask)
	r.PUT("/tasks/:id", tc.UpdateTask)
	r.DELETE("/tasks/:id", tc.DeleteTask)

	return r
}
