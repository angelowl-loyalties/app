package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"

	"github.com/spf13/viper"
)

type Config struct {
	DBConnString string `mapstructure:"DB_CONN_STRING"`
	DBKeyspace   string `mapstructure:"DB_KEYSPACE"`
	DBTable      string `mapstructure:"DB_TABLE"`
	Broker       string `mapstructure:"BROKER_HOST"`
	Topic        string `mapstructure:"TOPIC"`
}

type EtcdConfig = map[string]string

var ExclusionsEtcd EtcdConfig
var CampaignsEtcd EtcdConfig

func LoadConfig() (config Config, err error) {
	// create an etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://etcd:2379"}, // replace with your etcd endpoints
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx := context.Background()

	exclusionsWatchCh := cli.Watch(ctx, "exclusion", clientv3.WithPrefix())
	campaignsWatchCh := cli.Watch(ctx, "campaign", clientv3.WithPrefix())

	go handleWatchEvents(exclusionsWatchCh)
	go handleWatchEvents(campaignsWatchCh)

	viper.AddConfigPath("./config")
	viper.SetConfigName("dev.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		_ = viper.BindEnv("DB_CONN_STRING")
		_ = viper.BindEnv("DB_KEYSPACE")
		_ = viper.BindEnv("DB_TABLE")
		_ = viper.BindEnv("BROKER_HOST")
		_ = viper.BindEnv("TOPIC")

		err = viper.Unmarshal(&config)
		return
	}

	err = viper.Unmarshal(&config)

	return
}

func handleWatchEvents(watchCh clientv3.WatchChan) {
	for watchResp := range watchCh {
		for _, event := range watchResp.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				ExclusionsEtcd[string(event.Kv.Key)] = string(event.Kv.Value)
			case clientv3.EventTypeDelete:
				delete(ExclusionsEtcd, string(event.Kv.Key))
			}
		}
	}
}
