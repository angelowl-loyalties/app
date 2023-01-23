package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBConnString string `mapstructure:"DB_CONN_STRING"`
	DBKeyspace   string `mapstructure:"DB_KEYSPACE"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPass       string `mapstructure:"DB_PASS"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName(".dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		// production use
		_ = viper.BindEnv("PORT")
		_ = viper.BindEnv("DB_CONN_STRING")
		_ = viper.BindEnv("DB_KEYSPACE")
		_ = viper.BindEnv("DB_USER")
		_ = viper.BindEnv("DB_PASS")
		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
