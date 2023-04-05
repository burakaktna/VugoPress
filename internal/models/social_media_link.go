package models

import "time"

type SocialMediaLink struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Title     string    `json:"title" gorm:"size:255;not null" validate:"required"`
	Url       string    `json:"url" gorm:"size:255;not null" validate:"required,url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
