package url

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ironowl1907/url_shortener/internal/middleware"
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
	var ok bool
	var owner models.User

	// Extract user from context
	owner, ok = c.Keys["user"].(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read user"})
		return
	}

	// Bind and validate JSON input
	if err := c.ShouldBindJSON(&incomeURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	// Validate URL format
	if _, err := url.ParseRequestURI(incomeURL.OriginalURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format", "url": incomeURL.OriginalURL})
		return
	}

	// Set the owner ID
	incomeURL.Owner = owner.ID

	// Check URL reachability if not ignored
	if !incomeURL.IgnoreResponse {
		if err := validateURLReachability(incomeURL.OriginalURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "URL not reachable",
				"details": err.Error(),
				"url":     incomeURL.OriginalURL,
			})
			return
		}
	}

	// Create the URL entry
	createdURL, err := CreateURL(&incomeURL, h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create URL",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "URL created successfully",
		"data":    createdURL,
	})
}

// validateURLReachability checks if a URL is reachable
func validateURLReachability(urlStr string) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // 10 second timeout
	}

	// Make HEAD request (more efficient than GET)
	resp, err := client.Head(urlStr)
	if err != nil {
		return fmt.Errorf("failed to reach URL: %w", err)
	}
	defer resp.Body.Close() // Always close response body

	// Check if status code indicates success
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("URL returned status code: %d", resp.StatusCode)
	}

	return nil
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
	router.POST("/urls", middleware.RequireAuth, urlHandler.CreateURLHandler)
	router.GET("/urls", urlHandler.GetAllURLsHandler)
	router.GET("/urls/:id", urlHandler.GetURLByIDHandler)
	router.PUT("/urls/:id", urlHandler.UpdateURLHandler)
	router.DELETE("/urls/:id", urlHandler.DeleteURLHandler)
	router.GET("/:shortCode", urlHandler.RedirectByShortCodeHandler)
}
