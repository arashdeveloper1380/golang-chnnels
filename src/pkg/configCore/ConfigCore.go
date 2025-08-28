package configCore

import (
	"fiber-gorm-channel-ecommerce/src/application/config"
	"github.com/spf13/viper"
	"log"
)

var configurations config.Config

func ConfigurationGet() config.Config {
	return configurations
}

func ConfigurationSet() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("src/application/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading config data", err)
	}

	if err := viper.Unmarshal(&configurations); err != nil {
		log.Fatal("unable to decode into struct")
	}
}
