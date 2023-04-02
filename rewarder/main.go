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
	dbHost := c.DBConnString
	dbPort := c.DBPort
	dbKeyspace := c.DBKeyspace
	dbTable := c.DBTable
	dbUser := c.DBUser
	dbPass := c.DBPass
	dbUseSSL := c.DBUseSSL
	dbCreateIndex := c.DBCreateIndex
	models.InitDB(dbHost, dbPort, dbKeyspace, dbTable, dbUser, dbPass, dbUseSSL, dbCreateIndex)
	models.ConnectDB(dbHost, dbPort, dbUser, dbPass, dbKeyspace, dbUseSSL)

	// setup etcd connection
	etcdEndpoints := c.EtcdEndpoints
	etcdUsername := c.EtcdUsername
	etcdPassword := c.EtcdPassword
	internal.InitEtcdClient(etcdEndpoints, etcdUsername, etcdPassword)

	// For benchmarking purposes, comment this out for crappy performance
	// internal.WatchEtcd()

	// For benchmarking purposes, this is FOR crappy performance
	internal.RefreshFromEtcd()

	// Broker address and topic
	kafkaBroker := c.Broker
	consumeFromTopic := c.Topic

	internal.SaramaConsume(kafkaBroker, consumeFromTopic)
}
