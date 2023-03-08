package main

import (
	// "github.com/behzod-ph/ex-rater-bot/bot/command"
	"github.com/behzod-ph/ex-rater-bot/bot/config"
	"github.com/behzod-ph/ex-rater-bot/bot/handler"
	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {


	bot,updates := config.InitializeTelgramBot()

	for update := range updates {
		if update.CallbackQuery != nil {
			handler.CallbackQueryHandler(bot, update)
		} else if update.Message != nil {
			handler.MessageHandler(bot, update)
		}
	}
}


