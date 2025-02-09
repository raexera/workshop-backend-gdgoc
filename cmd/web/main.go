package main

import (
	"todo-list-service/infrastructure"
	"todo-list-service/internal/api"
)

func main() {
	// DB Connection
	db := infrastructure.NewDBConnection()
	defer db.Close()

	// HTTP
	routes := api.InitRoutes(db)
	routes.Run(":8080")
}
