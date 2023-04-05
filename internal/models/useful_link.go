package models

import "time"

type UsefulLink struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"size:255;not null;unique" validate:"required"`
	Url       string    `json:"url" gorm:"text;not null" validate:"required,url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
}
