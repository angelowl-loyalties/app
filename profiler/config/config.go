package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBConnString string `mapstructure:"DB_CONN_STRING"`
	AWSAccessKey string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	JWTKMSKey    string `mapstructure:"JWT_KMS_KEY_ID"`
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
		_ = viper.BindEnv("AWS_ACCESS_KEY_ID")
		_ = viper.BindEnv("AWS_SECRET_ACCESS_KEY")
		_ = viper.BindEnv("JWT_KMS_KEY_ID")
		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
