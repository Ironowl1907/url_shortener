package models

import "time"

type ShortenedURL struct {
	ID          uint      `gorm:"primaryKey"`
	OriginalURL string    `gorm:"type:text;not null"`
	ShortCode   string    `gorm:"type:varchar(16);unique;not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`

	OwnerID *uint `gorm:"index"`                         // Foreign key
	Owner   *User `gorm:"constraint:OnDelete:SET NULL;"` // Soft relation
}
