package url

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	fmt.Println("Init url routing")

	router.POST("/urls", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST /urls",
		})
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
