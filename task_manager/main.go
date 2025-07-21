package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	service := data.NewTaskService()
	r := router.SetupRouter(service)
	r.Run(":8080")
}
