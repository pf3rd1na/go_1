package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string `json:"key"`
}

func NewConfig() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	key, exists := os.LookupEnv("key")
	if !exists {
		fmt.Println("Environment variable 'key' not set")
	}
	fmt.Println("key:", key)
	return &Config{
		Key: key,
	}
}
