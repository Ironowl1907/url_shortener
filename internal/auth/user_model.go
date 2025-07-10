package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"not null" json:"user_name"`
	Email     string         `gorm:"not null;unique" json:"user_email"`
	Password  string         `gorm:"not null" json:"user_password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
