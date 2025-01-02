package handlers

func AskHandler(h *BotHandler, chatID int64, msg string) {
	if msg == "" {
		h.TelegramBot.SendMessage(chatID, "Provide /ask <question>")
		return
	} else {
		h.TelegramBot.SendMessage(chatID, msg)
	}
}
