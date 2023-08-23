package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {

	router := gin.Default()

	initializeRoutes(router)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	router.Run(":" + port)
}