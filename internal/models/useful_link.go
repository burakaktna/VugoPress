package models

import (
	"github.com/jinzhu/gorm"
)

type UsefulLink struct {
	gorm.Model
	Title  string `json:"title" gorm:"size:255;not null;unique" validate:"required"`
	Url    string `json:"url" gorm:"text;not null" validate:"required,url"`
	UserID uint   `json:"user_id" gorm:"index;not null"`
}
