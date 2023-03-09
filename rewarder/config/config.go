package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBConnString  string `mapstructure:"DB_CONN_STRING"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBKeyspace    string `mapstructure:"DB_KEYSPACE"`
	DBTable       string `mapstructure:"DB_TABLE"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	Broker        string `mapstructure:"BROKER_HOST"`
	Topic         string `mapstructure:"TOPIC"`
	EtcdEndpoints string `mapstructure:"ETCD_ENDPOINTS"`
	EtcdUsername  string `mapstructure:"ETCD_USERNAME"`
	EtcdPassword  string `mapstructure:"ETCD_PASSWORD"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName(".dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		_ = viper.BindEnv("DB_CONN_STRING")
		_ = viper.BindEnv("DB_PORT")
		_ = viper.BindEnv("DB_KEYSPACE")
		_ = viper.BindEnv("DB_TABLE")
		_ = viper.BindEnv("DB_USER")
		_ = viper.BindEnv("DB_PASS")
		_ = viper.BindEnv("BROKER_HOST")
		_ = viper.BindEnv("TOPIC")
		_ = viper.BindEnv("ETCD_ENDPOINTS")
		_ = viper.BindEnv("ETCD_USERNAME")
		_ = viper.BindEnv("ETCD_PASSWORD")

		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
