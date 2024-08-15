package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Port string
	//APIUrl string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := Config{
		DB: struct {
			Host     string
			Port     string
			User     string
			Password string
			Name     string
		}{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Port: os.Getenv("PORT"),
		//APIUrl: os.Getenv("API_URL"),
	}
	return &config
}
