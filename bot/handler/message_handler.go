package handler

import (
	// "github.com/behzod-ph/ex-rater-bot/bot/menu"
	"github.com/behzod-ph/ex-rater-bot/bot/menu"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := update.Message.Text
	switch {
	case text == "/start":
		StartCommandHandler(bot, update)
	}
}

const (
	StartCommand  = "/start"
	HelpCommand   = "/help"
	LogoutCommand = "/logout"
)

type BotCommandHandler interface {
	Handle()
}

func StartCommandHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please , choose langauge")
	msg.ReplyMarkup = menu.LanguageOption
	bot.Send(msg)
}
