package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type TagService interface {
	Create(tag *models.Tag) (*models.Tag, error)
	Index() ([]*models.Tag, error)
	Show(id uint) (*models.Tag, error)
	Update(id uint, updates *models.Tag) (*models.Tag, error)
	Delete(id uint) error
}

type tagService struct {
	db *gorm.DB
}

func NewTagService(db *gorm.DB) TagService {
	return &tagService{db: db}
}

func (s *tagService) Create(tag *models.Tag) (*models.Tag, error) {
	err := s.db.Create(tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *tagService) Index() ([]*models.Tag, error) {
	var tags []*models.Tag
	err := s.db.Find(&tags).Error
	return tags, err
}

func (s *tagService) Show(id uint) (*models.Tag, error) {
	var tag models.Tag
	err := s.db.First(&tag, id).Error
	return &tag, err
}

func (s *tagService) Update(id uint, updates *models.Tag) (*models.Tag, error) {
	var tag models.Tag
	if err := s.db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	err := s.db.Model(&tag).Updates(updates).Error
	return &tag, err
}

func (s *tagService) Delete(id uint) error {
	return s.db.Delete(&models.Tag{}, id).Error
}
