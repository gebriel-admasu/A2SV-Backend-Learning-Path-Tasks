package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Service *data.TaskService
}

func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{Service: service}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks := tc.Service.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := tc.Service.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	created := tc.Service.CreateTask(task)
	c.JSON(http.StatusCreated, created)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var updated models.Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	task, err := tc.Service.UpdateTask(id, updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	err = tc.Service.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
