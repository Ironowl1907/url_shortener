package url

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ironowl1907/url_shortener/internal/models"
	"gorm.io/gorm"
)

type URLHandler struct {
	DB *gorm.DB
}

// NewURLHandler creates a new URLHandler instance
func NewURLHandler(db *gorm.DB) *URLHandler {
	return &URLHandler{DB: db}
}

// CreateURLHandler handles POST /urls
func (h *URLHandler) CreateURLHandler(c *gin.Context) {
	var incomeURL models.URLPost
	if err := c.ShouldBindBodyWithJSON(&incomeURL); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if !incomeURL.IgnoreResponse {
		response, err := http.Get(incomeURL.OriginalURL)
		fmt.Println(incomeURL.IgnoreResponse)
		if err != nil || response.StatusCode != 200 {
			c.JSON(401, gin.H{"status": "warning, url not reachable", "url": incomeURL})
			return
		}
	}

	_, err := CreateURL(&incomeURL, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"status": "Server error", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "received", "url": incomeURL})
}

// GetAllURLsHandler handles GET /urls
func (h *URLHandler) GetAllURLsHandler(c *gin.Context) {
	var urls []models.ShortenedUrl
	h.DB.Find(&urls)
	c.JSON(200, urls)
}

// GetURLByIDHandler handles GET /urls/:id
func (h *URLHandler) GetURLByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var urls []models.ShortenedUrl
	err := h.DB.Raw("SELECT * FROM shortened_urls WHERE id = ?", id).Scan(&urls).Error
	if err != nil {
		c.JSON(500, gin.H{"status": "Server error", "error": err})
		return
	}
	c.JSON(200, urls)
}

// UpdateURLHandler handles PUT /urls/:id
func (h *URLHandler) UpdateURLHandler(c *gin.Context) {
	id := c.Param("id")
	var receivedURL models.ShortenedUrl
	var url models.ShortenedUrl

	// Bind JSON to struct
	if err := c.ShouldBindBodyWithJSON(&receivedURL); err != nil {
		c.JSON(400, gin.H{"status": "incorrect fields", "error": err.Error()})
		return
	}

	// Find existing URL record
	if err := h.DB.First(&url, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"status": "URL not found"})
			return
		}
		c.JSON(500, gin.H{"status": "Server error", "error": err.Error()})
		return
	}

	// Prepare updates
	updates := make(map[string]interface{})
	if receivedURL.OriginalURL != "" {
		updates["original_url"] = receivedURL.OriginalURL
	}
	if receivedURL.ShortCode != "" {
		if len(receivedURL.ShortCode) != 5 {
			c.JSON(400, gin.H{"status": "Request error", "error": "invalid shortened_url id"})
			return
		}
		updates["short_code"] = receivedURL.ShortCode
	}
	updates["updated_at"] = time.Now()

	// Apply updates
	if err := h.DB.Model(&url).Updates(updates).Error; err != nil {
		c.JSON(500, gin.H{"status": "Update failed", "error": err.Error()})
		return
	}

	// Fetch updated record
	if err := h.DB.First(&url, id).Error; err != nil {
		c.JSON(500, gin.H{"status": "Server error", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success", "data": url})
}

// DeleteURLHandler handles DELETE /urls/:id
func (h *URLHandler) DeleteURLHandler(c *gin.Context) {
	id := c.Param("id")
	err := h.DB.Delete(&models.ShortenedUrl{}, id).Error
	if err != nil {
		c.JSON(500, gin.H{"status": "Server error", "error": err})
		return
	}
	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})
}

// RedirectByShortCodeHandler handles GET /:shortCode
func (h *URLHandler) RedirectByShortCodeHandler(c *gin.Context) {
	shortCode := c.Param("shortCode")
	var result struct {
		OriginalURL string `gorm:"column:original_url"`
	}

	err := h.DB.Raw("SELECT original_url FROM shortened_urls WHERE short_code = ?",
		shortCode).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Short code not found for this URL")
			c.JSON(404, gin.H{
				"message": "Invalid code",
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
}

// Route sets up all URL routes using the extracted handlers
func Route(router *gin.Engine, dbConnection *gorm.DB) {
	fmt.Println("Init url routing")

	// Create handler instance
	urlHandler := NewURLHandler(dbConnection)

	// Register routes with extracted handlers
	router.POST("/urls", urlHandler.CreateURLHandler)
	router.GET("/urls", urlHandler.GetAllURLsHandler)
	router.GET("/urls/:id", urlHandler.GetURLByIDHandler)
	router.PUT("/urls/:id", urlHandler.UpdateURLHandler)
	router.DELETE("/urls/:id", urlHandler.DeleteURLHandler)
	router.GET("/:shortCode", urlHandler.RedirectByShortCodeHandler)
}
