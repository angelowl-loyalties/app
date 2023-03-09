package main

import (
	"log"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/config"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/internal"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// setup DB and connect
	dbConnString := c.DBConnString
	dbKeyspace := c.DBKeyspace
	dbTable := c.DBTable
	dbUser := c.DBUser
	dbPass := c.DBPass
	models.InitDB(dbConnString, dbKeyspace, dbTable, dbUser, dbPass)
	models.ConnectDB(dbConnString, dbKeyspace)

	// setup etcd connection
	etcdEndpoints := c.EtcdEndpoints
	etcdUsername := c.EtcdUsername
	etcdPassword := c.EtcdPassword
	internal.InitEtcdClient(etcdEndpoints, etcdUsername, etcdPassword)
	internal.WatchEtcd()

	// while loop to test etcd without consuming from kafka
	//for {}

	// Broker address and topic
	kafkaBroker := c.Broker
	consumeFromTopic := c.Topic

	internal.SaramaConsume(kafkaBroker, consumeFromTopic)
}
