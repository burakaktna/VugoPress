package services

import (
	"crypto/tls"
	"fmt"
	"github.com/burakaktna/VugoPress/internal/config"
	"io"
	"net/smtp"
)

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type emailService struct {
	cfg *config.Config
}

func NewEmailService(cfg *config.Config) EmailService {
	return &emailService{cfg: cfg}
}

func (s *emailService) SendEmail(to, subject, body string) error {
	from := s.cfg.SmtpUsername
	password := s.cfg.SmtpPassword
	smtpHost := s.cfg.SmtpHost
	smtpPort := s.cfg.SmtpPort
	auth := smtp.PlainAuth("", from, password, smtpHost)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", smtpHost, smtpPort), tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}

	defer func(client *smtp.Client) {
		err := client.Quit()
		if err != nil {
			panic(err)
		}
	}(client)

	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(from); err != nil {
		return err
	}

	if err = client.Rcpt(to); err != nil {
		return err
	}

	wc, err := client.Data()
	if err != nil {
		return err
	}

	defer func(wc io.WriteCloser) {
		err := wc.Close()
		if err != nil {
			panic(err)
		}
	}(wc)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	_, err = wc.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
}
