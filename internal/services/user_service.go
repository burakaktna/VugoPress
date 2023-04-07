package services

import (
	"fmt"
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.UserPost) (*models.UserDTO, error)
	Login(email string, password string) (string, error)
	CreateUser(user *models.UserPost) (*models.UserDTO, error)
	GetUsers() ([]*models.UserDTO, error)
	GetUser(id uint) (*models.UserDTO, error)
	UpdateUser(id uint, updates *models.UserPost) (*models.UserDTO, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo   repository.UserRepository
	appKey string
}

func NewUserService(repo repository.UserRepository, appKey string) UserService {
	return &userService{
		repo:   repo,
		appKey: appKey,
	}
}

func (s *userService) Register(userPost *models.UserPost) (*models.UserDTO, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPost.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := userPost.ToUser()
	user.Password = string(hashedPassword)
	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !user.CheckPassword(password) {
		return "", fmt.Errorf("incorrect password")
	}

	token, err := user.GenerateToken(s.appKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) CreateUser(userPost *models.UserPost) (*models.UserDTO, error) {
	user := userPost.ToUser()
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *userService) GetUsers() ([]*models.UserDTO, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := make([]*models.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = user.ToDTO()
	}

	return userDTOs, nil
}

func (s *userService) GetUser(id uint) (*models.UserDTO, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *userService) UpdateUser(id uint, updates *models.UserPost) (*models.UserDTO, error) {
	updatedUser, err := s.repo.UpdateUser(id, updates.ToUser())
	if err != nil {
		return nil, err
	}

	return updatedUser.ToDTO(), nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
