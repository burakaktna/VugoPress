package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/repository"
)

type SocialMediaLinkService interface {
	CreateSocialMediaLink(socialMediaLink *models.SocialMediaLink) (*models.SocialMediaLink, error)
	GetSocialMediaLinks() ([]*models.SocialMediaLink, error)
	GetSocialMediaLink(id uint) (*models.SocialMediaLink, error)
	UpdateSocialMediaLink(id uint, updates *models.SocialMediaLink) (*models.SocialMediaLink, error)
	DeleteSocialMediaLink(id uint) error
}

type socialMediaLinkService struct {
	repo repository.SocialMediaLinkRepository
}

func NewSocialMediaLinkService(repo repository.SocialMediaLinkRepository) SocialMediaLinkService {
	return &socialMediaLinkService{repo: repo}
}

func (s *socialMediaLinkService) CreateSocialMediaLink(socialMediaLink *models.SocialMediaLink) (*models.SocialMediaLink, error) {
	err := s.repo.CreateSocialMediaLink(socialMediaLink)
	if err != nil {
		return nil, err
	}

	return socialMediaLink, nil
}

func (s *socialMediaLinkService) GetSocialMediaLinks() ([]*models.SocialMediaLink, error) {
	return s.repo.GetSocialMediaLinks()
}

func (s *socialMediaLinkService) GetSocialMediaLink(id uint) (*models.SocialMediaLink, error) {
	return s.repo.GetSocialMediaLink(id)
}

func (s *socialMediaLinkService) UpdateSocialMediaLink(id uint, updates *models.SocialMediaLink) (*models.SocialMediaLink, error) {
	return s.repo.UpdateSocialMediaLink(id, updates)
}

func (s *socialMediaLinkService) DeleteSocialMediaLink(id uint) error {
	return s.repo.DeleteSocialMediaLink(id)
}
