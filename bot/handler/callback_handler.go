package handler

import (
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CallbackQueryHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update)  {
	callbackData := update.CallbackQuery.Data
	switch  {
	case strings.Contains(callbackData, "lang") : LanguageQeuryHandler(bot,update)
	
}
}

func LanguageQeuryHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update){
	
}

