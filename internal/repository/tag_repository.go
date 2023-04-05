package repository

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type TagRepository interface {
	CreateTag(tag *models.Tag) error
	GetTags() ([]*models.Tag, error)
	GetTag(id uint) (*models.Tag, error)
	UpdateTag(id uint, updates *models.Tag) (*models.Tag, error)
	DeleteTag(id uint) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) CreateTag(tag *models.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) GetTags() ([]*models.Tag, error) {
	var tags []*models.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *tagRepository) GetTag(id uint) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.Where("id = ?", id).First(&tag).Error
	return &tag, err
}

func (r *tagRepository) UpdateTag(id uint, updates *models.Tag) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	tag.Title = updates.Title
	tag.Url = updates.Url
	err = r.db.Save(&tag).Error
	return &tag, err
}

func (r *tagRepository) DeleteTag(id uint) error {
	return r.db.Delete(&models.Tag{}, "id = ?", id).Error
}
