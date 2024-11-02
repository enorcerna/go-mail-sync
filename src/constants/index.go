package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AuthMailConfig struct {
	HostName string
	Email    string
	Password string
}

func newAuthMainConfig() *AuthMailConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed .env file")
	}
	host := os.Getenv("HOST")
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	if host == "" || email == "" || password == "" {
		log.Println("WARNING: Is email,host,password environment not value ")
	}
	return &AuthMailConfig{
		HostName: host,
		Email:    email,
		Password: password,
	}
}

var AuthMail = newAuthMainConfig()
