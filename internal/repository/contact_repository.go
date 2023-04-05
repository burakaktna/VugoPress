package repository

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type ContactRepository interface {
	CreateContact(contact *models.Contact) error
	GetContacts() ([]*models.Contact, error)
	GetContact(id uint) (*models.Contact, error)
	UpdateContact(id uint, updates *models.Contact) (*models.Contact, error)
	DeleteContact(id uint) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) CreateContact(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *contactRepository) GetContacts() ([]*models.Contact, error) {
	var contacts []*models.Contact
	err := r.db.Find(&contacts).Error
	return contacts, err
}

func (r *contactRepository) GetContact(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.Where("id = ?", id).First(&contact).Error
	return &contact, err
}

func (r *contactRepository) UpdateContact(id uint, updates *models.Contact) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.Where("id = ?", id).First(&contact).Error
	if err != nil {
		return nil, err
	}
	contact.Name = updates.Name
	contact.Email = updates.Email
	contact.Message = updates.Message
	err = r.db.Save(&contact).Error
	return &contact, err
}

func (r *contactRepository) DeleteContact(id uint) error {
	return r.db.Delete(&models.Contact{}, "id = ?", id).Error
}
