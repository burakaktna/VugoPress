package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type UsefulLinkService interface {
	Create(usefulLink *models.UsefulLink) (*models.UsefulLink, error)
	Index() ([]*models.UsefulLink, error)
	Show(id uint) (*models.UsefulLink, error)
	Update(id uint, updates *models.UsefulLink) (*models.UsefulLink, error)
	Delete(id uint) error
}

type usefulLinkService struct {
	db *gorm.DB
}

func NewUsefulLinkService(db *gorm.DB) UsefulLinkService {
	return &usefulLinkService{db: db}
}

func (s *usefulLinkService) Create(usefulLink *models.UsefulLink) (*models.UsefulLink, error) {
	err := s.db.Create(usefulLink).Error
	if err != nil {
		return nil, err
	}
	return usefulLink, nil
}

func (s *usefulLinkService) Index() ([]*models.UsefulLink, error) {
	var links []*models.UsefulLink
	err := s.db.Find(&links).Error
	return links, err
}

func (s *usefulLinkService) Show(id uint) (*models.UsefulLink, error) {
	var link models.UsefulLink
	err := s.db.First(&link, id).Error
	return &link, err
}

func (s *usefulLinkService) Update(id uint, updates *models.UsefulLink) (*models.UsefulLink, error) {
	var link models.UsefulLink
	if err := s.db.First(&link, id).Error; err != nil {
		return nil, err
	}
	err := s.db.Model(&link).Updates(updates).Error
	return &link, err
}

func (s *usefulLinkService) Delete(id uint) error {
	return s.db.Delete(&models.UsefulLink{}, id).Error
}
