package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Broker string `mapstructure:"BROKER_HOST"`
	Topic  string `mapstructure:"TOPIC"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		// production use
		_ = viper.BindEnv("BROKER_HOST")
		_ = viper.BindEnv("TOPIC")
		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
