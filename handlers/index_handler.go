package handlers

import (
	"github.com/azisuverse/jirothebot/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	TelegramBot   *services.TelegramBot
	OpenAIService *services.OpenAIService
}

func NewBotHandler(token string, apiKey string) (*BotHandler, error) {
	telegramBot := services.NewTelegramBot(token)
	openAIService := services.NewOpenAIService(apiKey)

	return &BotHandler{
		TelegramBot:   telegramBot,
		OpenAIService: openAIService,
	}, nil
}

func (h *BotHandler) UpdateMsgChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := h.TelegramBot.Bot.GetUpdatesChan(u)

	return updates
}

func (h *BotHandler) Handlers() {
	updates := h.UpdateMsgChannel()

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "start":
			StartHandler(h, update.Message.Chat.ID)
		case "help":
			HelpHandler(h, update.Message.Chat.ID)
		case "ask":
			query := update.Message.CommandArguments()
			if query == "" {
				AskHandler(h, update.Message.Chat.ID, "")
				break
			}

			resp, err := h.OpenAIService.GetResponse(query)
			if err != nil {
				h.TelegramBot.SendMessage(update.Message.Chat.ID, "Error: "+err.Error())
				continue
			}

			AskHandler(h, update.Message.Chat.ID, resp)
		case "author":
			AuthorHandler(h, update.Message.Chat.ID)
		case "follow":
			FollowHandler(h, update.Message.Chat.ID)
		default:
			h.TelegramBot.SendMessage(update.Message.Chat.ID, "Unknown command. Use /ask <question>")
		}
	}
}
