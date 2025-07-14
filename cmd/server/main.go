package main

import (
	"log"
	"os"

	"github.com/ironowl1907/url_shortener/internal/app"
	"github.com/ironowl1907/url_shortener/internal/tools"
)

func main() {
	// Load .env file
	{
		err := tools.LoadEnvVariables()
		if err != nil {
			log.Panic("error loading .env file: " + err.Error())
		}
	}

	// Create new app
	app := app.NewApp()

	// Initialize the DB
	if err := app.InitDB(); err != nil {
		panic("Couldn't connect to database")
	}

	// Setup the routes
	app.SetupRoutes()

	// Run the server
	if err := app.Run(os.Getenv("PORT")); err != nil {
		panic("Couldn't run server")
	}
}
