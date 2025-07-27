package data

import (
	"context"
	"errors"
	"task_manager_with_auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func getUserCollection() *mongo.Collection {
	return GetCollection("users")
}

// CreateUser registers a new user
func CreateUser(username, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if username already exists
	count, _ := getUserCollection().CountDocuments(ctx, bson.M{"username": username})
	if count > 0 {
		return models.User{}, errors.New("username already exists")
	}

	// First user is admin
	role := "user"
	total, _ := getUserCollection().CountDocuments(ctx, bson.M{})
	if total == 0 {
		role = "admin"
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	newUser := models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	_, err := getUserCollection().InsertOne(ctx, newUser)
	if err != nil {
		return models.User{}, err
	}

	newUser.Password = "" // Hide password in response
	return newUser, nil
}

// AuthenticateUser validates username and password
func AuthenticateUser(username, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := getUserCollection().FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("invalid username or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return models.User{}, errors.New("invalid username or password")
	}
	return user, nil
}

// PromoteUser updates a user's role to admin
func PromoteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := getUserCollection().UpdateOne(ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{"role": "admin"}})

	if err != nil || res.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
