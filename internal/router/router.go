package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ironowl1907/url_shortener/internal/auth"
	"github.com/ironowl1907/url_shortener/internal/url"
	"github.com/ironowl1907/url_shortener/internal/user"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, dbConnection *gorm.DB) {
	auth.Route(router, dbConnection)
	user.Route(router, dbConnection)
	url.Route(router, dbConnection)
}
