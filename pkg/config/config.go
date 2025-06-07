package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DB_DSN    string
	JWTSecret string
}

func LoadEnv() *Config {
	err := godotenv.Load()
		if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		DB_DSN:    os.Getenv("DB_DSN"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}