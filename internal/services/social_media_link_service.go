package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type SocialMediaLinkService interface {
	Create(socialMediaLink *models.SocialMediaLink) (*models.SocialMediaLink, error)
	Index() ([]*models.SocialMediaLink, error)
	Show(id uint) (*models.SocialMediaLink, error)
	Update(id uint, updates *models.SocialMediaLink) (*models.SocialMediaLink, error)
	Delete(id uint) error
}

type socialMediaLinkService struct {
	db *gorm.DB
}

func NewSocialMediaLinkService(db *gorm.DB) SocialMediaLinkService {
	return &socialMediaLinkService{db: db}
}

func (s *socialMediaLinkService) Create(socialMediaLink *models.SocialMediaLink) (*models.SocialMediaLink, error) {
	err := s.db.Create(socialMediaLink).Error
	if err != nil {
		return nil, err
	}
	return socialMediaLink, nil
}

func (s *socialMediaLinkService) Index() ([]*models.SocialMediaLink, error) {
	var links []*models.SocialMediaLink
	err := s.db.Find(&links).Error
	return links, err
}

func (s *socialMediaLinkService) Show(id uint) (*models.SocialMediaLink, error) {
	var link models.SocialMediaLink
	err := s.db.First(&link, id).Error
	return &link, err
}

func (s *socialMediaLinkService) Update(id uint, updates *models.SocialMediaLink) (*models.SocialMediaLink, error) {
	var link models.SocialMediaLink
	if err := s.db.First(&link, id).Error; err != nil {
		return nil, err
	}
	err := s.db.Model(&link).Updates(updates).Error
	return &link, err
}

func (s *socialMediaLinkService) Delete(id uint) error {
	return s.db.Delete(&models.SocialMediaLink{}, id).Error
}
