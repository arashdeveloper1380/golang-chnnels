package configCore

import (
	"fiber-gorm-channel-ecommerce/application/config"
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
	viper.AddConfigPath("application/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading config data")
	}

	if err := viper.Unmarshal(&configurations); err != nil {
		log.Fatal("unable to decode into struct")
	}
}
