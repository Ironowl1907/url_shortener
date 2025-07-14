package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ironowl1907/url_shortener/internal/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func RequireAuth(c *gin.Context) {
	// Get the cookie from Req
	tokenString, err := c.Cookie("JWT")
	if err != nil {
		c.AbortWithError(403, err)
		return
	}

	// Decode and validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithError(403, err)
		}

		// Find user with token sub
		var user models.User
		response := DB.First(&user, claims["sub"].(float64))
		if response.Error != nil {
			c.AbortWithError(403, response.Error)
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.AbortWithError(403, err)
	}
}
