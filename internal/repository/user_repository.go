package repository

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetUser(id uint) (*models.User, error)
	UpdateUser(id uint, updates *models.User) (*models.User, error)
	DeleteUser(id uint) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUser(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) UpdateUser(id uint, updates *models.User) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.SiteName = updates.SiteName
	user.Email = updates.Email
	user.SiteDomain = updates.SiteDomain
	user.Name = updates.Name
	user.Surname = updates.Surname
	user.Phone = updates.Phone
	user.Address = updates.Address

	err = r.db.Save(&user).Error
	return &user, err
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
