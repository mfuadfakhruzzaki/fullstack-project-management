// config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    Port         string
    DBHost       string
    DBPort       string
    DBUser       string
    DBPassword   string
    DBName       string
    JWTSecret    string
    SMTPHost     string
    SMTPPort     string
    SMTPUser     string
    SMTPPassword string
}

var AppConfig Config

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    AppConfig = Config{
        Port:         os.Getenv("PORT"),
        DBHost:       os.Getenv("DB_HOST"),
        DBPort:       os.Getenv("DB_PORT"),
        DBUser:       os.Getenv("DB_USER"),
        DBPassword:   os.Getenv("DB_PASSWORD"),
        DBName:       os.Getenv("DB_NAME"),
        JWTSecret:    os.Getenv("JWT_SECRET"),
        SMTPHost:     os.Getenv("SMTP_HOST"),
        SMTPPort:     os.Getenv("SMTP_PORT"),
        SMTPUser:     os.Getenv("SMTP_USER"),
        SMTPPassword: os.Getenv("SMTP_PASSWORD"),
    }
}
