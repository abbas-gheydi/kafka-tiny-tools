package kafkat

import (
	"fmt"
	"log"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TopicInfo(topics []string, kafkaServer string, groupName string, autooffsetrest string, sessionTimeOut int) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  kafkaServer,
		"group.id":           groupName,
		"auto.offset.reset":  autooffsetrest,
		"session.timeout.ms": sessionTimeOut,
	})

	if err != nil {
		panic(err)
	}

	m, e := c.GetMetadata(&topics[0], true, 30000)
	if e != nil {
		log.Println(e)
	} else {
		fmt.Println("brokers:", m.Brokers)

		for key := range m.Topics {
			if !strings.HasPrefix(key, "_") {
				fmt.Println("topic name:", key)

			}
		}
	}

	c.Close()

}
