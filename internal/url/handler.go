package url

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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
			return
		}

		c.JSON(200, gin.H{"status": "received", "url": url})
	})

	router.GET("/urls", func(c *gin.Context) {
		var urls []ShortenedUrl
		dbConnection.Find(&urls)
		c.JSON(200, urls)
	})

	router.GET("/urls/:id", func(c *gin.Context) {
		id := c.Param("id")
		var urls []ShortenedUrl
		err := dbConnection.Raw("SELECT * FROM shortened_urls WHERE id = ?", id).Scan(&urls).Error
		if err != nil {
			c.JSON(500, gin.H{"status": "Server error", "error": err})
			return
		}
		c.JSON(200, urls)
	})

	router.PUT("/urls/:id", func(c *gin.Context) {
		id := c.Param("id")
		var recievedURL ShortenedUrl
		var url ShortenedUrl
		{
			err := c.ShouldBindBodyWithJSON(&recievedURL)
			if err != nil {
				c.JSON(400, gin.H{"status": "incorrect fields", "error": err.Error()})
				return
			}
		}
		{
			err := dbConnection.First(&url, id).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(404, gin.H{"status": "URL not found"})
					return
				}
				c.JSON(500, gin.H{"status": "Server error", "error": err.Error()})
				return
			}
		}

		updates := make(map[string]interface{})

		if recievedURL.OriginalURL != "" {
			updates["original_url"] = recievedURL.OriginalURL
		}
		if recievedURL.ShortCode != "" {
			if len(recievedURL.ShortCode) != 5 {
				c.JSON(400, gin.H{"status": "Request error", "error": "invalid shortened_url id"})
				return
			}
			updates["short_code"] = recievedURL.ShortCode
		}

		updates["updated_at"] = time.Now()

		{
			err := dbConnection.Model(&url).Updates(updates).Error
			if err != nil {
				c.JSON(500, gin.H{"status": "Update failed", "error": err.Error()})
				return
			}
		}

		{
			err := dbConnection.First(&url, id).Error
			if err != nil {
				c.JSON(500, gin.H{"status": "Server error", "error": err.Error()})
				return
			}
		}

		c.JSON(200, gin.H{"status": "success", "data": url})
	})
	router.DELETE("/urls/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := dbConnection.Delete(&ShortenedUrl{}, id).Error
		if err != nil {
			c.JSON(500, gin.H{"status": "Server error", "error": err})
			return
		}

		c.JSON(200, gin.H{
			"message": "deleted succesfully",
		})
	})
	router.GET("/:shortCode", func(c *gin.Context) {
		id := c.Param("shortCode")

		var result struct {
			OriginalURL string `gorm:"column:original_url"`
		}

		err := dbConnection.Raw("SELECT original_url FROM shortened_urls WHERE short_code = ?",
			id).First(&result).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("Short code not found for this URL")
				c.JSON(404, gin.H{
					"message": "Code not found",
				})
			} else {
				log.Printf("Database error: %v", err)
				c.JSON(500, gin.H{
					"message": "Internal server error",
				})
			}
			return
		}

		c.Redirect(http.StatusMovedPermanently, result.OriginalURL)
	})
}
