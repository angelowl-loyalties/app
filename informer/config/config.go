package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBConnString  string `mapstructure:"DB_CONN_STRING"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBKeyspace    string `mapstructure:"DB_KEYSPACE"`
	DBTable       string `mapstructure:"DB_TABLE"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBUseSSL      bool   `mapstructure:"DB_SSL"`
	DBCreateIndex bool   `mapstructure:"DB_CREATE_INDEX"`
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
		_ = viper.BindEnv("DB_PORT")
		_ = viper.BindEnv("DB_KEYSPACE")
		_ = viper.BindEnv("DB_TABLE")
		_ = viper.BindEnv("DB_USER")
		_ = viper.BindEnv("DB_PASS")
		_ = viper.BindEnv("DB_SSL")
		_ = viper.BindEnv("DB_CREATE_INDEX")
		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
