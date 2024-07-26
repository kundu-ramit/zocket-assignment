package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	database "github.com/kundu-ramit/zocket/infra/database"
	"github.com/kundu-ramit/zocket/routes"
)

func main() {
	// Get the command-line arguments
	args := os.Args[1:]

	// Check the number of arguments
	if len(args) == 0 {
		log.Fatal("No command specified.")
	}

	// Handle the command
	switch args[0] {
	case "server":
		startServer()
	default:
		log.Fatal("Invalid command:", args[0])
	}
}

func startServer() {
	router := gin.Default()
	database.Initialize()
	routes.RegisterRoutes(router)
	router.Run(":8002")
}
