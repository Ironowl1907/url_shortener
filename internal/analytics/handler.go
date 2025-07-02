package analytics

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	fmt.Println("Init analytics routing")

	router.GET("/urls/:id/stats", func(c *gin.Context) {
		// id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "GET /urls/:id/stats",
		})
	})
	router.GET("/stats/overview", func(c *gin.Context) {
		// id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "GET /stats/overview",
		})
	})
	router.GET("/urls/:id/stats", func(c *gin.Context) {
		// id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "GET /urls/:id/stats",
		})
	})
}
