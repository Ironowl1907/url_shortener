package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ironowl1907/url_shortener/internal/middleware"
	"github.com/ironowl1907/url_shortener/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	// Get email and password from json
	var body struct {
		Email    string
		Name     string
		password string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed to read body",
		})
		return
	}

	// Verify password
	if len(body.password) < 8 {
		c.JSON(400, gin.H{
			"status": "password should be at least 8 characters",
		})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "failed to hash password",
		})
		return
	}
	// Create the user
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

	result := h.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"status": "failed to create user",
		})
		return
	}

	// Respond
	c.JSON(200, gin.H{
		"status": "ok",
		"user":   user,
	})
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	// Get creadentials from body
	var body struct {
		Email    string
		password string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed to read body",
		})
		return
	}
	// DB lookup
	var user models.User
	response := h.DB.First(&user, "email = ?", body.Email)
	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"status": "invalid email or password",
		})
		return
	} else if response.Error != nil {
		c.JSON(500, gin.H{
			"status": "failed to read database",
		})
		return
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.password))
	if err != nil {
		c.JSON(400, gin.H{
			"status": "invalid email or passoword",
		})
		return
	}

	// Generate JWT token

	key := os.Getenv("SECRET")
	if key == "" {
		log.Fatal("failed to retrieve secret from .env file")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		c.JSON(500, gin.H{
			"status": "failed to create JWT",
		})
		log.Println(err.Error())
		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("JWT", s, 3600*24, "", "", false, true)
	c.JSON(200, gin.H{})
}

func (h *AuthHandler) ValidateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logged In",
	})
}

func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "POST /auth/logout",
	})
}

func (h *AuthHandler) GetMeHandler(c *gin.Context) {
	// Get user from context
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		c.JSON(400, gin.H{
			"status": "failed to get user from context",
		})
		return
	}
	// Respond
	c.JSON(200, user)
}

func (h *AuthHandler) UpdateMeHandler(c *gin.Context) {
	// Get user from context
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		c.JSON(400, gin.H{
			"status": "failed to get user from context",
		})
		return
	}

	// Parse json
	var body struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed to read body",
			"error":  err.Error(),
		})
		return
	}

	// Validate input
	if body.Email == "" && body.Name == "" && body.Password == "" {
		c.JSON(400, gin.H{
			"status": "no fields provided for update",
		})
		return
	}

	// Email validation (basic)
	if body.Email != "" {
		if !isValidEmail(body.Email) {
			c.JSON(400, gin.H{
				"status": "invalid email format",
			})
			return
		}
	}

	// Password validation
	if body.Password != "" {
		if len(body.Password) < 8 {
			c.JSON(400, gin.H{
				"status": "password must be at least 8 characters",
			})
			return
		}
	}

	// Update fields
	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Name != "" {
		user.Name = body.Name
	}
	if body.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "failed to hash password",
			})
			return
		}
		user.Password = string(hash)
	}

	// DB submit
	response := h.DB.Save(&user)
	if response.Error != nil {
		c.JSON(500, gin.H{
			"status": "failed to update database",
			"error":  response.Error.Error(),
		})
		return
	}

	userResponse := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	c.JSON(200, gin.H{
		"status": "user updated successfully",
		"user":   userResponse,
	})
}

// Helper function for email validation
func isValidEmail(email string) bool {
	// Should user a more robust solution
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func (h *AuthHandler) DeleteMeHandler(c *gin.Context) {
}

func Route(router *gin.Engine, dbConnection *gorm.DB) {
	fmt.Println("Init auth routing")

	authHandler := NewAuthHandler(dbConnection)

	router.POST("/auth/register", authHandler.RegisterHandler)
	router.POST("/auth/login", authHandler.LoginHandler)
	router.POST("/auth/logout", middleware.RequireAuth, authHandler.LogoutHandler)
	router.GET("/auth/validate", middleware.RequireAuth, authHandler.ValidateUser)
	router.GET("/auth/me", middleware.RequireAuth, authHandler.GetMeHandler)
	router.PUT("/auth/me", middleware.RequireAuth, authHandler.UpdateMeHandler)
	router.DELETE("/auth/me", middleware.RequireAuth, authHandler.DeleteMeHandler)
}
