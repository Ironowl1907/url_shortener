package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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
	// Get User from context
	user := c.Keys["user"]
	// Respond
	c.JSON(200, user)
}

func (h *AuthHandler) UpdateMeHandler(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "PUT /auth/me",
	})
}

func (h *AuthHandler) DeleteMeHandler(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "DELETE /auth/me",
	})
}

func Route(router *gin.Engine, dbConnection *gorm.DB) {
	fmt.Println("Init auth routing")

	authHandler := NewAuthHandler(dbConnection)

	router.POST("/auth/register", authHandler.RegisterHandler)
	router.POST("/auth/login", authHandler.LoginHandler)
	router.POST("/auth/logout", authHandler.LogoutHandler)
	router.GET("/auth/validate", middleware.RequireAuth, authHandler.ValidateUser)
	router.GET("/auth/me", middleware.RequireAuth, authHandler.GetMeHandler)
	router.PUT("/auth/me", middleware.RequireAuth, authHandler.UpdateMeHandler)
	router.DELETE("/auth/me", authHandler.DeleteMeHandler)
}
