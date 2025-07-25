package controllers

import (
	"net/http"
	"task_manager_with_DB/data"
	"task_manager_with_DB/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService *data.TaskService
}

func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{TaskService: service}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.TaskService.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.TaskService.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.TaskService.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.TaskService.UpdateTask(id, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.TaskService.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
