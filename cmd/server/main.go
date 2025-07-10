package main

import (
	"github.com/ironowl1907/url_shortener/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.InitDB(); err != nil {
		panic("Couldn't connect to database")
	}
	app.SetupRoutes()

	if err := app.Run("8080"); err != nil {
		panic("Couldn't run server")
	}
}
