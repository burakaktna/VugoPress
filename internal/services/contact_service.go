package services

import (
	"fmt"
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
	"log"
)

type ContactService interface {
	Create(contact *models.Contact) (*models.Contact, error)
	Index() ([]*models.Contact, error)
	Show(id uint) (*models.Contact, error)
	Update(updates *models.Contact) (*models.Contact, error)
	Delete(id uint) error
}

type contactService struct {
	db           *gorm.DB
	emailService EmailService
}

func NewContactService(db *gorm.DB, emailService EmailService) ContactService {
	return &contactService{
		db:           db,
		emailService: emailService,
	}
}

func (s *contactService) Create(contact *models.Contact) (*models.Contact, error) {
	err := s.db.Create(contact).Error
	if err != nil {
		return nil, err
	}

	go sendContactEmail(contact, err, s)

	return contact, nil
}

func (s *contactService) Index() ([]*models.Contact, error) {
	var contacts []*models.Contact
	err := s.db.Find(&contacts).Error
	return contacts, err
}

func (s *contactService) Show(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := s.db.First(&contact, id).Error
	return &contact, err
}

func (s *contactService) Update(updates *models.Contact) (*models.Contact, error) {
	return updates, s.db.Updates(&updates).Error
}

func (s *contactService) Delete(id uint) error {
	return s.db.Delete(&models.Contact{}, id).Error
}

func sendContactEmail(contact *models.Contact, err error, s *contactService) {
	to := "aktunamuhasebe04@gmail.com"
	subject := "Yeni İletişim Formu"
	body := fmt.Sprintf("İsim: %s\nE-Posta: %s\n\nTelefon Numarası: %s\nMesaj: %s", contact.Name, contact.Email, contact.Phone, contact.Message)

	err = s.emailService.SendEmail(to, subject, body)
	if err != nil {
		log.Printf("E-posta gönderilemedi: %v", err)
	}
}
