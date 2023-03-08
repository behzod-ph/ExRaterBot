package config

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitializeTelgramBot() (*tgbotapi.BotAPI ,tgbotapi.UpdatesChannel){
	botConfig, err := LoadBotConfig("./")



	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(botConfig.SecretAccessToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = botConfig.Timeouts

	updates := bot.GetUpdatesChan(u)
	return bot,updates;
}