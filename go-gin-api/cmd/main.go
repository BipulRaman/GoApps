package main

import (
	"go-gin-api/internal/routes"
)

func main() {
	// Initialize the router
	r := routes.SetupRouter()

	// Start the server
	//r.Run(":8080")
	r.Run()
}
