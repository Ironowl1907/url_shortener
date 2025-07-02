package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	fmt.Println("Init auth routing")

	router.POST("/auth/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST /auth/register",
		})
	})
	router.POST("/auth/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST /auth/login",
		})
	})
	router.POST("/auth/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST /auth/logout",
		})
	})
	router.GET("/auth/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET /auth/me",
		})
	})
	router.PUT("/auth/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT /auth/me",
		})
	})
	router.DELETE("/auth/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE /auth/me",
		})
	})
}
