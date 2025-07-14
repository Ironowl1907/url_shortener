package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ironowl1907/url_shortener/internal/auth"
	"github.com/ironowl1907/url_shortener/internal/url"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, dbConnection *gorm.DB) {
	auth.Route(router, dbConnection)
	url.Route(router, dbConnection)
}
