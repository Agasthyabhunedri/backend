package main

import (
	"fmt"
	"go-rest-api/config"

	"go-rest-api/models"
	"go-rest-api/router"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	config.ConnectDatabase()
	// Migrate the schema
	config.DB.AutoMigrate(&models.User{})
	r := router.NewRouter()
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type"}))(r)))
}
