package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/repository"
)

type TagService interface {
	CreateTag(tag *models.Tag) (*models.Tag, error)
	GetTags() ([]*models.Tag, error)
	GetTag(id uint) (*models.Tag, error)
	UpdateTag(id uint, updates *models.Tag) (*models.Tag, error)
	DeleteTag(id uint) error
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{repo: repo}
}

func (s *tagService) CreateTag(tag *models.Tag) (*models.Tag, error) {
	err := s.repo.CreateTag(tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *tagService) GetTags() ([]*models.Tag, error) {
	return s.repo.GetTags()
}

func (s *tagService) GetTag(id uint) (*models.Tag, error) {
	return s.repo.GetTag(id)
}

func (s *tagService) UpdateTag(id uint, updates *models.Tag) (*models.Tag, error) {
	return s.repo.UpdateTag(id, updates)
}

func (s *tagService) DeleteTag(id uint) error {
	return s.repo.DeleteTag(id)
}
