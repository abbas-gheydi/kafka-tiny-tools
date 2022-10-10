package kafkat

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer(topics []string, kafkaServer string, groupName string, autooffsetrest string, sessionTimeOut int) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  kafkaServer,
		"group.id":           groupName,
		"auto.offset.reset":  autooffsetrest,
		"session.timeout.ms": sessionTimeOut,
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(topics, nil)
	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
