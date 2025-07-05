package url

import (
	"time"

	"github.com/ironowl1907/url_shortener/internal/user"
)

type ShortenedURL struct {
	ID          uint      `gorm:"primaryKey"`
	OriginalURL string    `gorm:"type:text;not null"`
	ShortCode   string    `gorm:"type:varchar(16);unique;not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`

	OwnerID *uint      `gorm:"index"`                         // Foreign key
	Owner   *user.User `gorm:"constraint:OnDelete:SET NULL;"` // Soft relation
}
