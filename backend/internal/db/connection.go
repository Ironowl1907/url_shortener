package db

import (
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB")
	if dsn == "" {
		return nil, errors.New("Database credentials undefined !")
	}
	dbConection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("couldn't open db")
	}

	return dbConection, nil
}
