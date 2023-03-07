package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBConnString  string `mapstructure:"DB_CONN_STRING"`
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
		fmt.Println(err)
		// production use
		_ = viper.BindEnv("PORT")
		_ = viper.BindEnv("DB_CONN_STRING")
		_ = viper.BindEnv("ETCD_ENDPOINTS")
		_ = viper.BindEnv("ETCD_USERNAME")
		_ = viper.BindEnv("ETCD_PASSWORD")
		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
