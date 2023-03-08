package config

import (
	"log"

	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

type BotConfig struct {
	BotUsername       string `mapstructure:"BOT_USERNAME"`
	SecretAccessToken string `mapstructure:"BOT_ACCESS_TOKEN"`
	Timeouts          int    `mapstructure:"BOT_TIMEOUTS"`
	UpdateOffset      int    `mapstructure:"BOT_UPDATE_OFFSET"`
}

type DataBaseConfig struct {
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Path     string `mapstructure:"DB_PATH"`
}
type ExchageRatesConfig struct {
	URL string `mapstructure:"EXCHANGE_RATES_URL"`
}


func LoadBotConfig(path string) (config BotConfig, err error) {
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

func LoadDataBaseConfig(path string) (config DataBaseConfig, err error) {
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

func LoadExchageRatesConfig(path string) (config ExchageRatesConfig, err error) {
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






