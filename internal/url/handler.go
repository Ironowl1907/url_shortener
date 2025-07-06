package url

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(router *gin.Engine, dbConnection *gorm.DB) {
	fmt.Println("Init url routing")

	router.POST("/urls", func(c *gin.Context) {
		var url URLPost
		if err := c.ShouldBindBodyWithJSON(&url); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		_, err := CreateURL(&url, dbConnection)
		if err != nil {
			c.JSON(500, gin.H{"status": "Server error", "url": url})
		}

		c.JSON(200, gin.H{"status": "received", "url": url})
	})

	router.GET("/urls", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET /urls",
		})
	})
	router.GET("/urls:id", func(c *gin.Context) {
		// id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "GET /urls:id",
		})
	})
	router.PUT("/urls:id", func(c *gin.Context) {
		// id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "PUT /urls:id",
		})
	})
	router.DELETE("/urls:id", func(c *gin.Context) {
		// id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "DELETE /urls:id",
		})
	})
	router.GET("/:shortCode", func(c *gin.Context) {
		// id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "GET /:shortCode",
		})
	})
}
