package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ironowl1907/url_shortener/internal/db"
	"github.com/ironowl1907/url_shortener/internal/router"
)

func main() {
	// Connect to the local instance of postgres
	_ = db.InitDB()

	// Routing with gin
	ginRouter := gin.Default()

	// Simple example for a endpoint
	ginRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Application Routing
	router.InitRouting(ginRouter)

	// Run the server
	ginRouter.Run()
}
