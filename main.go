package main

import (
	"fmt"
	"go-rest-api/config"

	"go-rest-api/models"
	"go-rest-api/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/handlers"
)

func main() {

	// ✅ Load environment variables from .env file (in dev only)
	_ = godotenv.Load()

	config.ConnectDatabase()
	// Migrate the schema
	config.DB.AutoMigrate(&models.User{})
	r := router.NewRouter()

	// Get port from environment (Render injects it)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	// ✅ Load CORS origins from env
	allowedOrigins := []string{}
	if origin := os.Getenv("CORS_ORIGIN_LOCAL"); origin != "" {
		allowedOrigins = append(allowedOrigins, origin)
	}
	if origin := os.Getenv("CORS_ORIGIN_PROD"); origin != "" {
		allowedOrigins = append(allowedOrigins, origin)
	}

	// CORS settings — allow your deployed frontend domain too
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Start server
	fmt.Println("Server is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
