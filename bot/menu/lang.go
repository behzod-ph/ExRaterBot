package menu

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"


var LanguageOption = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("πΊπΏ Uzbek","lang_uz"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("πΊπΈ English","lang_eng"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("π·πΊ Russian","lang_rus"),
	),
)
