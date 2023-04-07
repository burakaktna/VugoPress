package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title" gorm:"size:255;not null" validate:"required"`
	Content string `json:"content" gorm:"type:text;not null" validate:"required"`
	UserID  uint   `json:"user_id" gorm:"index;not null"`
	Tags    []*Tag `json:"tags" gorm:"many2many:article_tags;constraint:OnUpdate:CASCADE;"`
}

type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
}

func (ArticleTag) TableName() string {
	return "article_tags"
}
