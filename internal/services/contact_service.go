package services

import (
	"fmt"
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/repository"
	"log"
)

type ContactService interface {
	CreateContact(contact *models.Contact) (*models.Contact, error)
	GetContacts() ([]*models.Contact, error)
	GetContact(id uint) (*models.Contact, error)
	UpdateContact(id uint, updates *models.Contact) (*models.Contact, error)
	DeleteContact(id uint) error
}

type contactService struct {
	repo         repository.ContactRepository
	emailService EmailService
}

func NewContactService(repo repository.ContactRepository, emailService EmailService) ContactService {
	return &contactService{
		repo:         repo,
		emailService: emailService,
	}
}

func (s *contactService) CreateContact(contact *models.Contact) (*models.Contact, error) {
	err := s.repo.CreateContact(contact)
	if err != nil {
		return nil, err
	}

	sendContactEmail(contact, err, s)

	return contact, nil
}

func (s *contactService) GetContacts() ([]*models.Contact, error) {
	return s.repo.GetContacts()
}

func (s *contactService) GetContact(id uint) (*models.Contact, error) {
	return s.repo.GetContact(id)
}

func (s *contactService) UpdateContact(id uint, updates *models.Contact) (*models.Contact, error) {
	return s.repo.UpdateContact(id, updates)
}

func (s *contactService) DeleteContact(id uint) error {
	return s.repo.DeleteContact(id)
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
