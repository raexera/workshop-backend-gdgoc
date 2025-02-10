package main

import (
	"os"

	"todo-list-service/infrastructure"
	"todo-list-service/internal/api"
)

func main() {
	// DB Connection
	db := infrastructure.NewDBConnection()
	defer db.Close()

	// HTTP Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes := api.InitRoutes(db)
	routes.Run(":8080")
}
