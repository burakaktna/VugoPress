package repository

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type SocialMediaLinkRepository interface {
	CreateSocialMediaLink(socialMediaLink *models.SocialMediaLink) error
	GetSocialMediaLinks() ([]*models.SocialMediaLink, error)
	GetSocialMediaLink(id uint) (*models.SocialMediaLink, error)
	UpdateSocialMediaLink(id uint, updates *models.SocialMediaLink) (*models.SocialMediaLink, error)
	DeleteSocialMediaLink(id uint) error
}

type socialMediaLinkRepository struct {
	db *gorm.DB
}

func NewSocialMediaLinkRepository(db *gorm.DB) SocialMediaLinkRepository {
	return &socialMediaLinkRepository{db: db}
}

func (r *socialMediaLinkRepository) CreateSocialMediaLink(socialMediaLink *models.SocialMediaLink) error {
	return r.db.Create(socialMediaLink).Error
}

func (r *socialMediaLinkRepository) GetSocialMediaLinks() ([]*models.SocialMediaLink, error) {
	var socialMediaLinks []*models.SocialMediaLink
	err := r.db.Find(&socialMediaLinks).Error
	return socialMediaLinks, err
}

func (r *socialMediaLinkRepository) GetSocialMediaLink(id uint) (*models.SocialMediaLink, error) {
	var socialMediaLink models.SocialMediaLink
	err := r.db.Where("id = ?", id).First(&socialMediaLink).Error
	return &socialMediaLink, err
}

func (r *socialMediaLinkRepository) UpdateSocialMediaLink(id uint, updates *models.SocialMediaLink) (*models.SocialMediaLink, error) {
	var socialMediaLink models.SocialMediaLink
	err := r.db.Where("id = ?", id).First(&socialMediaLink).Error
	if err != nil {
		return nil, err
	}
	socialMediaLink.Title = updates.Title
	socialMediaLink.Url = updates.Url
	err = r.db.Save(&socialMediaLink).Error
	return &socialMediaLink, err
}

func (r *socialMediaLinkRepository) DeleteSocialMediaLink(id uint) error {
	return r.db.Delete(&models.SocialMediaLink{}, "id = ?", id).Error
}
