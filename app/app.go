package app

import (
	"fmt"

	"github.com/Amangoyal1A/studentapi/configs"
	"github.com/Amangoyal1A/studentapi/routes"
	"github.com/Amangoyal1A/studentapi/storage"
	"github.com/gin-gonic/gin"
)

// Start initializes and starts the application
func Start() {
	// Attempt to connect to the database
	dbConn, err := storage.Connect()
	if err != nil {
		fmt.Println("Failed to connect to the database: ", err)
	}
	fmt.Println("Database connected successfully", dbConn)

	configs.AutoMigrate(dbConn)

	// Initialize Gin router
	router := gin.Default()

	// Pass the DB instance to the routes
	routes.StudentRoutes(router, dbConn)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
