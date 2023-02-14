package main

import (
	"fmt"
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

	// Broker address and topic
	kafkaBroker := c.Broker
	consumeFromTopic := c.Topic

	var transaction = models.Seed_transaction
	fmt.Println(transaction)
	internal.ProcessMessage(transaction)

	internal.SaramaConsume(kafkaBroker, consumeFromTopic)
}
