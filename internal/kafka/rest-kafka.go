package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/paroar/roadtracing-rest-kafka/internal/types"
)

var producer *kafka.Producer

var (
	kafkaServer = os.Getenv("KAFKA_SERVER")
	topic       = os.Getenv("KAFKA_TOPIC")
)

// NewProducer creates a new kafka producer
func NewProducer() {

	for {
		fmt.Println("Creating Producer")
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": kafkaServer,
		})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Producer created")
			producer = p
			break
		}
	}

}

// SavePositionToKafka saves the received position to kafka server
func SavePositionToKafka(pos types.Position) {

	jsonString, err := json.Marshal(pos)
	if err != nil {
		log.Println(err)
	}

	posString := string(jsonString)

	log.Printf("KAFKA_PRODUCER: %s", posString)

	for _, word := range []string{string(posString)} {
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(word),
		}, nil)
		if err != nil {
			log.Println(err)
		}
	}

	producer.Flush(100)

}
