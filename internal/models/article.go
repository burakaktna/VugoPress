package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"size:255;not null" validate:"required"`
	Content   string    `json:"content" gorm:"type:text;not null" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	Tags      []*Tag    `json:"tags" gorm:"many2many:article_tags;constraint:OnUpdate:CASCADE;"`
}

type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
}

func (ArticleTag) TableName() string {
	return "article_tags"
}
