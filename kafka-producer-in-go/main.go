package main

import (
	"log"

	kafka_config "kafka_implementation.com/producer"
)

func main() {

	// Initialize the Kafka producer
	producer, err := kafka_config.NewKafkaProducer()
	if err != nil {
		log.Fatal("producer is not initialized: ", err)
	}
	// Publish a message to Kafka
	producer.Publish()
	log.Println("Messages published successfully")
}
