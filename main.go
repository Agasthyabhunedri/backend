package main

import (
	"fmt"
	"go-rest-api/config"

	"go-rest-api/models"
	"go-rest-api/router"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	config.ConnectDatabase()
	// Migrate the schema
	config.DB.AutoMigrate(&models.User{})
	r := router.NewRouter()
	fmt.Println("Server is running on http://localhost:8080")
	// Get port from environment (Render injects it)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	// CORS settings — allow your deployed frontend domain too
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",                    // for local dev
			"https://frontend-eta-beryl-37.vercel.app", // replace with your actual Vercel URL
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)

	// Start server
	fmt.Println("Server is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
