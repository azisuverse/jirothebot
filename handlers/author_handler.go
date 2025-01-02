package handlers

func AuthorHandler(h *BotHandler, chatID int64) {
  h.TelegramBot.SendMessage(
    chatID,
    `
    Author Profile
Username - @showmegoods
Bio - ONE MORE TIME
Website - showmegoods.vercel.app
Follow - /follow
    `,
    )
}
