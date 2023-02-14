package main

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/internal"
)

func main() {
	// create a new context
	// ctx := context.Background()

	dbConnString := c.DBConnString
	dbKeyspace := c.DBKeyspace
	dbTable := c.DBTable
	// dbUser := c.DBUser
	// dbPass := c.DBPass
	models.InitDB(dbConnString, dbKeyspace, dbTable)
	models.ConnectDB(dbConnString, dbKeyspace)

	// Broker address and topic
	kafkaBroker := c.Broker
	consumeFromTopic := c.Topic

	// internal.Consume(ctx, c.Broker)
	internal.SaramaConsume()
}
