package main

import (
	"fmt"
	"log"
	"net/http"
	"restaurant-backend/database"
	"restaurant-backend/routes"
)

func main() {
	// Connect to database
	database.Connect()

	// Setup routes
	routes.UserRoutes()
	routes.OrderRoutes()

	// Start the server
	fmt.Println("Server is running on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
