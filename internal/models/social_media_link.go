package models

import (
	"github.com/jinzhu/gorm"
)

type SocialMediaLink struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"not null"`
	Title  string `json:"title" gorm:"size:255;not null" validate:"required"`
	Url    string `json:"url" gorm:"size:255;not null" validate:"required,url"`
}
