package app

import (
	"fmt"

	"github.com/Amangoyal1A/studentapi/storage"
)

// Start initializes and starts the application
func Start() {
	// Attempt to connect to the database
	dbConn, err := storage.Connect()
	if err != nil {
		fmt.Println("Failed to connect to the database: ", err)
	}
	fmt.Println("Database connected successfully", dbConn)

}
