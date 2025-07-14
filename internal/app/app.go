package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ironowl1907/url_shortener/internal/db"
	"github.com/ironowl1907/url_shortener/internal/middleware"
	"github.com/ironowl1907/url_shortener/internal/router"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewApp() *App {
	return &App{
		DB:     nil,
		Router: gin.Default(),
	}
}

func (app *App) InitDB() error {
	var err error
	app.DB, err = db.InitDB()
	if err != nil {
		return err
	}
	middleware.SetDB(app.DB)
	return nil
}

func (app *App) Run(port string) error {
	routes := app.Router.Routes()
	fmt.Println("Registered routes:")
	for _, route := range routes {
		fmt.Printf("%s %s --> %s\n", route.Method, route.Path, route.Handler)
	}

	// Run the server
	if err := app.Router.Run(":" + port); err != nil {
		return err
	}
	return nil
}

func (app *App) SetupRoutes() error {
	app.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.SetupRoutes(app.Router, app.DB)
	return nil
}
