package services

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot *tgbotapi.BotAPI
}

func NewTelegramBot(token string) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Failed to initialize Bot: %v", err)
	}

	log.Printf("Bot authorized on account: %s", bot.Self.UserName)

	return &TelegramBot{
		Bot: bot,
	}
}

func (tgbot *TelegramBot) SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)

	_, err := tgbot.Bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
