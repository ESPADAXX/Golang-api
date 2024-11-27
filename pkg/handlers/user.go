package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go-chi-app.com/go-chi-app/pkg/config"
	"go-chi-app.com/go-chi-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch all users from the MongoDB collection
	cursor, err := config.MongoDB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		http.Error(w, "Failed to parse users", http.StatusInternalServerError)
		return
	}

	// Respond with the list of users
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the MongoDB collection
	result, err := config.MongoDB.Collection("users").InsertOne(ctx, bson.M{
		"name": user.Name, // Insert only the "name" field; MongoDB will generate "_id"
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Extract the inserted ID and convert it to a string
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	// Respond with the created user including the generated ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.User{
		ID:   insertedID,
		Name: user.Name,
	})
}
