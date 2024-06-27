package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: 0,
	})
	fmt.Print("Step 3")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("could not read message %v", err)
			break
		}
		fmt.Printf("received: %s\n", string(msg.Value))
	}
}
