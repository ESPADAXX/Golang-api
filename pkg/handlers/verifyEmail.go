package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go-chi-app.com/go-chi-app/pkg/config"
	
)
var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the query
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Verification token is required", http.StatusBadRequest)
		return
	}

	// Parse and validate the token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	email := claims["email"].(string)

	// Update the user's verification status in the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := config.MongoDB.Collection("users").UpdateOne(ctx, bson.M{"email": email}, bson.M{
		"$set": bson.M{"verified": true},
	})
	if err != nil || result.ModifiedCount == 0 {
		http.Error(w, "Failed to verify email", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Email verified successfully. You can now log in.",
	})
}
