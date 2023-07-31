package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
} // sebagai wadah kosong untuk menyimpan informasi tertentu

func NewEnvironment() *Environment {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return &Environment{}
}

func (e *Environment) Get(key string) string {
	return os.Getenv(key)
}
