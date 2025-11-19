package tools

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file: " + err.Error())
		return err
	}

	return nil
}
