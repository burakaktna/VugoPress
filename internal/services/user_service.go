package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUser(id uint) (*models.User, error)
	UpdateUser(id uint, updates *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUsers() ([]*models.User, error) {
	return s.repo.GetUsers()
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *userService) UpdateUser(id uint, updates *models.User) (*models.User, error) {
	return s.repo.UpdateUser(id, updates)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
