package handlers

func FollowHandler(h *BotHandler, chatID int64) {
	h.TelegramBot.SendMessage(
		chatID,
		`
    Folllow the Author
Telegram - t.me/showmegoods
GitHub - github.com/showmegoods
X/Twitter - x.com/showmegoods
    `,
	)
}
