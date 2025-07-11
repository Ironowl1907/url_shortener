package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	user := User{Name: body.Name, Email: body.Email, Password: string(hash)}

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
	c.JSON(501, gin.H{
		"message": "POST /auth/login",
	})
}

func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "POST /auth/logout",
	})
}

func (h *AuthHandler) GetMeHandler(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "GET /auth/me",
	})
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
	router.GET("/auth/me", authHandler.GetMeHandler)
	router.PUT("/auth/me", authHandler.UpdateMeHandler)
	router.DELETE("/auth/me", authHandler.DeleteMeHandler)
}
