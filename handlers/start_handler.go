package handlers

func StartHandler(h *BotHandler, chatID int64) {
  h.TelegramBot.SendMessage(
    chatID,
    `
    Welcome to JiroTheBot!
type /help to show the available commands.
    `,
    )
}
