package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go-chi-app.com/go-chi-app/pkg/config"
	"go-chi-app.com/go-chi-app/pkg/routes"
)
func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}
	
	// Initialize Google OAuth2 configuration
	config.InitOAuthConfig()

}
func main() {
	// Connect to MongoDB
	mongoURI := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	config.ConnectMongoDB(mongoURI, dbName)

	// Set up the router
	router := routes.SetupRouter()

	// Start the HTTP server
	log.Println("Server is running on http://localhost:5050")
	err := http.ListenAndServe(":5050", router)
	if err != nil {
		log.Fatal(err)
	}
}
