package main

import (
	"log"

	"Training/session-2-latihan-crud-user-gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Routes
	router.SetupRouter(r)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
