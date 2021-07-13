package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strings"
)


func NewConsumer(ctx context.Context) error {
	brokers := strings.Split("localhost:9092", ",")

	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         "g1",
		Topic:           "newTopic",
		//MinBytes:        10e3,            // 10KB
		//MaxBytes:        10e6,            // 10MB
		//MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		//ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println(err)
			continue
		}


		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
	}
}
