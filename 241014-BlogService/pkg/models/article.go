package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	AuthorID uint   `gorm:"not null"`
}
