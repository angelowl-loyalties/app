package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBConnString  string `mapstructure:"DB_CONN_STRING"`
	DBKeyspace    string `mapstructure:"DB_KEYSPACE"`
	DBTable       string `mapstructure:"DB_TABLE"`
	Broker        string `mapstructure:"BROKER_HOST"`
	Topic         string `mapstructure:"TOPIC"`
	EtcdEndpoints string `mapstructure:"ETCD_ENDPOINTS"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName(".dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		_ = viper.BindEnv("DB_CONN_STRING")
		_ = viper.BindEnv("DB_KEYSPACE")
		_ = viper.BindEnv("DB_TABLE")
		_ = viper.BindEnv("BROKER_HOST")
		_ = viper.BindEnv("TOPIC")
		_ = viper.BindEnv("ETCD_ENDPOINTS")

		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
