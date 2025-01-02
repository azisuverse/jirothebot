package main

import (
	"github.com/azisuverse/jirothebot/config"
	"github.com/azisuverse/jirothebot/handlers"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	bot, err := handlers.NewBotHandler(cfg.TelegramBotToken, cfg.OpenAIAPIKey)

	if err != nil {
		log.Fatal("Failed to initialize Bot: ", err)
	}

	bot.Handlers()
}
