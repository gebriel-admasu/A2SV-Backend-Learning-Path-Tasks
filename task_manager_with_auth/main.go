package main

import (
	"task_manager_with_auth/data"
	"task_manager_with_auth/router"
)

func main() {
	data.InitMongoDB() // Must run before accessing collections
	r := router.SetupRouter()
	r.Run(":8080")
}
