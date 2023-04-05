package repository

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type UsefulLinkRepository interface {
	CreateUsefulLink(usefulLink *models.UsefulLink) error
	GetUsefulLinks() ([]*models.UsefulLink, error)
	GetUsefulLink(id uint) (*models.UsefulLink, error)
	UpdateUsefulLink(id uint, updates *models.UsefulLink) (*models.UsefulLink, error)
	DeleteUsefulLink(id uint) error
}

type usefulLinkRepository struct {
	db *gorm.DB
}

func NewUsefulLinkRepository(db *gorm.DB) UsefulLinkRepository {
	return &usefulLinkRepository{db: db}
}

func (r *usefulLinkRepository) CreateUsefulLink(usefulLink *models.UsefulLink) error {
	return r.db.Create(usefulLink).Error
}

func (r *usefulLinkRepository) GetUsefulLinks() ([]*models.UsefulLink, error) {
	var usefulLinks []*models.UsefulLink
	err := r.db.Find(&usefulLinks).Error
	return usefulLinks, err
}

func (r *usefulLinkRepository) GetUsefulLink(id uint) (*models.UsefulLink, error) {
	var usefulLink models.UsefulLink
	err := r.db.Where("id = ?", id).First(&usefulLink).Error
	return &usefulLink, err
}

func (r *usefulLinkRepository) UpdateUsefulLink(id uint, updates *models.UsefulLink) (*models.UsefulLink, error) {
	var usefulLink models.UsefulLink
	err := r.db.Where("id = ?", id).First(&usefulLink).Error
	if err != nil {
		return nil, err
	}
	usefulLink.Title = updates.Title
	usefulLink.Url = updates.Url
	err = r.db.Save(&usefulLink).Error
	return &usefulLink, err
}

func (r *usefulLinkRepository) DeleteUsefulLink(id uint) error {
	return r.db.Delete(&models.UsefulLink{}, "id = ?", id).Error
}
