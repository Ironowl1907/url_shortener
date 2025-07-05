package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ironowl1907/url_shortener/internal/auth"
	"github.com/ironowl1907/url_shortener/internal/url"
	"github.com/ironowl1907/url_shortener/internal/user"
)

func SetupRoutes(router *gin.Engine) {
	auth.Route(router)
	user.Route(router)
	url.Route(router)
}
