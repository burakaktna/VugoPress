package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/repository"
)

type UsefulLinksService interface {
	CreateUsefulLink(usefulLink *models.UsefulLink) (*models.UsefulLink, error)
	GetUsefulLinks() ([]*models.UsefulLink, error)
	GetUsefulLink(id uint) (*models.UsefulLink, error)
	UpdateUsefulLink(id uint, updates *models.UsefulLink) (*models.UsefulLink, error)
	DeleteUsefulLink(id uint) error
}

type usefulLinksService struct {
	repo repository.UsefulLinkRepository
}

func NewUsefulLinksService(repo repository.UsefulLinkRepository) UsefulLinksService {
	return &usefulLinksService{repo: repo}
}

func (s *usefulLinksService) CreateUsefulLink(usefulLink *models.UsefulLink) (*models.UsefulLink, error) {
	err := s.repo.CreateUsefulLink(usefulLink)
	if err != nil {
		return nil, err
	}

	return usefulLink, nil
}

func (s *usefulLinksService) GetUsefulLinks() ([]*models.UsefulLink, error) {
	return s.repo.GetUsefulLinks()
}

func (s *usefulLinksService) GetUsefulLink(id uint) (*models.UsefulLink, error) {
	return s.repo.GetUsefulLink(id)
}

func (s *usefulLinksService) UpdateUsefulLink(id uint, updates *models.UsefulLink) (*models.UsefulLink, error) {
	return s.repo.UpdateUsefulLink(id, updates)
}

func (s *usefulLinksService) DeleteUsefulLink(id uint) error {
	return s.repo.DeleteUsefulLink(id)
}
