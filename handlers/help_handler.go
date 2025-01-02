package handlers

func HelpHandler(h *BotHandler, chatID int64) {
  h.TelegramBot.SendMessage(
    chatID,
    `
    Help Menu
/start - Starting JiroTheBot
/help - Show the help menu
/ask  - Ask the question
/author - Show the author profile
/follow - Follow the author
    `,
    )
}
