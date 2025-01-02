package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
	OpenAIAPIKey     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.")
	}

	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		OpenAIAPIKey:     os.Getenv("OPENAI_API_KEY"),
	}
}
