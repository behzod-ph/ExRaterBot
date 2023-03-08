package util

import (
	"log"

	viper "github.com/spf13/viper"
	"github.com/behzod-ph/ex-rater-bot/bot/config"
)


func LoadBotConfig(path string) (config config.BotConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	return
}
