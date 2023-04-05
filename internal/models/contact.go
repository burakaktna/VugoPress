package models

import "time"

type Contact struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"size:255;not null" validate:"required"`
	Email     string    `json:"email" gorm:"size:255;not null" validate:"required,email"`
	Phone     string    `json:"phone" gorm:"size:255;not null" validate:"required"`
	Message   string    `json:"message" gorm:"type:text;not null" validate:"required"`
	Processed bool      `json:"processed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
}
