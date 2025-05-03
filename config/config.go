package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	OpenAIAPIKey      string
	JavaBackendAuthURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{
		Port:              os.Getenv("PORT"),
		OpenAIAPIKey:      os.Getenv("OPENAI_API_KEY"),
		JavaBackendAuthURL: os.Getenv("JAVA_BACKEND_AUTH_URL"),
	}

	return config, nil
}
