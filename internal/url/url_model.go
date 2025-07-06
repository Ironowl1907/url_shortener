package url

import (
	"time"
)

type ShortenedUrl struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	OriginalURL string    `gorm:"not null" json:"original_url"`
	ShortCode   string    `gorm:"unique;not null" json:"short_code"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// UserOwner   *user.User `gorm:"foreignKey:Owner;references:ID;constraint:OnDelete:SET NULL"` // Relationship
}

type URLPost struct {
	OriginalURL string `json:"url"`
	Owner       *uint  `json:"owner"`
}
