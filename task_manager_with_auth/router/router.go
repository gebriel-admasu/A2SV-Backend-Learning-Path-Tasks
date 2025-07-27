package router

import (
	"task_manager_with_auth/controllers"
	"task_manager_with_auth/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/tasks", controllers.GetAllTasks)
	r.GET("/tasks/:id", controllers.GetTaskByID)

	// Protected routes (JWT required)
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware()) // Require JWT
	{
		admin.PUT("/promote/:username", middleware.AdminOnly(), controllers.PromoteUser)
		admin.POST("/tasks", middleware.AdminOnly(), controllers.CreateTask)
		admin.PUT("/tasks/:id", middleware.AdminOnly(), controllers.UpdateTask)
		admin.DELETE("/tasks/:id", middleware.AdminOnly(), controllers.DeleteTask)
	}

	return r
}
