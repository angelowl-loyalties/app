package main

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/internal"
)

func main() {
	// create a new context
	// ctx := context.Background()

	// c, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatalln("Failed at config", err)
	// }

	// internal.Consume(ctx, c.Broker)
	internal.SaramaConsume()
}
