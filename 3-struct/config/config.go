package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string `json:"key"`
}

func NewConfig() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		Key: os.Getenv("key"),
	}
}
