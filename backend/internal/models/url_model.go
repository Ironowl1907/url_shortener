package models

import (
	"time"
)

type ShortenedUrl struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	OriginalURL string    `gorm:"not null" json:"original_url"`
	ShortCode   string    `gorm:"unique;not null" json:"short_code"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Title       string    `gorm:"default:untitled" json:"title"`
	Description string    `gorm:"default:" json:"description"`
	OwnerID     uint      `gorm:"not null" json:"owner_id"`
	Owner       User      `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
}

type URLPost struct {
	OriginalURL    string `json:"url"`
	Owner          uint   `json:"-"`
	IgnoreResponse bool   `json:"ignore_response"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}
