package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppKey           string
	AppPort          string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	SmtpHost         string
	SmtpPort         string
	SmtpUsername     string
	SmtpPassword     string
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Hata: .env dosyası yüklenemedi. %v", err)
	}

	return &Config{
		AppKey:           os.Getenv("APP_KEY"),
		AppPort:          os.Getenv("APP_PORT"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		SmtpHost:         os.Getenv("SMTP_HOST"),
		SmtpPort:         os.Getenv("SMTP_PORT"),
		SmtpUsername:     os.Getenv("SMTP_USERNAME"),
		SmtpPassword:     os.Getenv("SMTP_PASSWORD"),
	}
}
