package services

import (
	"fmt"

	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.UserPost) (*models.UserDTO, error)
	Login(email string, password string) (string, error)
	Create(user *models.UserPost) (*models.UserDTO, error)
	Index() ([]*models.UserDTO, error)
	Show(id uint) (*models.UserDTO, error)
	Update(id uint, updates *models.UserPost) (*models.UserDTO, error)
	Delete(id uint) error
}

type userService struct {
	db     *gorm.DB
	appKey string
}

func NewUserService(db *gorm.DB, appKey string) UserService {
	return &userService{
		db:     db,
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
	err = s.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *userService) Login(email, password string) (string, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
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

func (s *userService) Create(userPost *models.UserPost) (*models.UserDTO, error) {
	user := userPost.ToUser()
	err := s.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *userService) Index() ([]*models.UserDTO, error) {
	var users []*models.User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	userDTOs := make([]*models.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = user.ToDTO()
	}

	return userDTOs, nil
}

func (s *userService) Show(id uint) (*models.UserDTO, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *userService) Update(id uint, updates *models.UserPost) (*models.UserDTO, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	updatedUser := updates.ToUser()
	err := s.db.Model(&user).Updates(updatedUser).Error
	if err != nil {
		return nil, err
	}
	return user.ToDTO(), nil
}

func (s *userService) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}
