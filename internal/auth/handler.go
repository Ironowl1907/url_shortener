package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "POST /auth/register",
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
