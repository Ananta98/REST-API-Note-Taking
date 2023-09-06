package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID  uint   `gorm:"not null;" json:"user_id"`
	Title   string `gorm:"size:255" json:"title"`
	Content string `gorm:"text" json:"content"`

	// Relationship
	user User `json:"-"`
}
