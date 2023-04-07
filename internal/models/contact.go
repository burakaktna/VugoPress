package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name      string `json:"name" gorm:"size:255;not null" validate:"required"`
	Email     string `json:"email" gorm:"size:255;not null" validate:"required,email"`
	Phone     string `json:"phone" gorm:"size:255;not null" validate:"required"`
	Message   string `json:"message" gorm:"type:text;not null" validate:"required"`
	Processed bool   `json:"processed" gorm:"default:false"`
	UserID    uint   `json:"user_id" gorm:"index;not null"`
}
