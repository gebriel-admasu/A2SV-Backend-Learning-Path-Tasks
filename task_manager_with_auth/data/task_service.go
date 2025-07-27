package data

import (
	"context"
	"task_manager_with_auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getTaskCollection() *mongo.Collection {
	return GetCollection("tasks")
}

// GetAllTasks retrieves all tasks
func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := getTaskCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var t models.Task
		cursor.Decode(&t)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// GetTaskByID retrieves a task by its ID
func GetTaskByID(id int) (models.Task, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task models.Task
	err := getTaskCollection().FindOne(ctx, bson.M{"id": id}).Decode(&task)
	return task, err == nil
}

// CreateTask creates a new task
func CreateTask(title, description string) (models.Task, error) {
	newTask := models.Task{
		ID:          int(primitive.NewObjectID().Timestamp().Unix()),
		Title:       title,
		Description: description,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := getTaskCollection().InsertOne(ctx, newTask)
	return newTask, err
}

// UpdateTask updates an existing task
func UpdateTask(id int, title, description string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := getTaskCollection().UpdateOne(ctx, bson.M{"id": id},
		bson.M{"$set": bson.M{"title": title, "description": description}})
	return err == nil && res.MatchedCount > 0
}

// DeleteTask removes a task by its ID
func DeleteTask(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := getTaskCollection().DeleteOne(ctx, bson.M{"id": id})
	return err == nil && res.DeletedCount > 0
}
