package models

import (
	"time"
)

type User struct {
	ID          uint         `json:"id" gorm:"primary_key"`
	SiteName    string       `json:"site_name" gorm:"size:255;not null" validate:"required"`
	Phone       string       `json:"phone" gorm:"size:20"`
	Email       string       `json:"email" gorm:"size:255"`
	Address     string       `json:"address" gorm:"type:text"`
	Name        string       `json:"name" gorm:"size:255;not null" validate:"required"`
	Surname     string       `json:"surname" gorm:"size:255;not null" validate:"required"`
	Password    string       `json:"-" gorm:"size:255;not null" validate:"required"`
	SiteDomain  string       `json:"site_domain" gorm:"size:255;not null" validate:"required"`
	Contacts    []Contact    `json:"contacts" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Articles    []Article    `json:"articles" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags        []Tag        `json:"tags" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UsefulLinks []UsefulLink `json:"useful_links" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}
