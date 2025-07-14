package url

import (
	"log"
	"math/rand"

	"github.com/ironowl1907/url_shortener/internal/models"
	"gorm.io/gorm"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	res := make([]rune, n)

	for i := range n {
		res[i] = letters[rand.Intn(len(letters))]
	}

	return string(res)
}

func CreateURL(url *models.URLPost, dbConnection *gorm.DB) (*models.ShortenedUrl, error) {
	shortenedURL := &models.ShortenedUrl{
		OriginalURL: url.OriginalURL,
		ShortCode:   RandSeq(5),
	}

	result := dbConnection.Create(shortenedURL)
	if result.Error != nil {
		log.Printf("Error creating shortened url: %v", result.Error)
		return nil, result.Error
	}

	log.Printf("Rows affected: %d\n", result.RowsAffected)

	return shortenedURL, nil
}
