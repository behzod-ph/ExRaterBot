package menu

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"


var LanguageOption = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇺🇿 Uzbek","lang_uz"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇺🇸 English","lang_eng"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Russian","lang_rus"),
	),
)
