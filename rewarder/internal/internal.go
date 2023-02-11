package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, brokerAddress string) {
	i := 0
	// l := log.New(os.Stdout, "kafka reader: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "transaction6",
		GroupID: "rewarder",
		MaxWait: time.Second,
		// MinBytes: 10e3,
		// MaxBytes: 10e6,
		// Logger:  l,
	})

	start := time.Now()

	for {
		_, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// fmt.Println("received: ", string(msg.Value))
		i++
		if i%1000 == 0 && i != 0 {
			end := time.Now()
			fmt.Println(end.Sub(start))
			start = time.Now()
		}
	}
}
