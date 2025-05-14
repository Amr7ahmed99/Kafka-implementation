// kafka/producer.go
package kafka

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	config "kafka_implementation.com/config"

	"github.com/IBM/sarama"
)

// KafkaProducer handles message publishing to Kafka
type KafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

// NewKafkaProducer initializes the Kafka sync producer
func NewKafkaProducer() (*KafkaProducer, error) {
	config := config.NewKafkaConfig()
	broker := getEnv("KAFKA_HOST", "localhost:9092")
	topic := getEnv("KAFKA_TOPIC", "my-kafka-topic")

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	return &KafkaProducer{producer: producer, topic: topic}, nil
}

// Publish sends a new message every second
func (kp *KafkaProducer) Publish() {
	defer kp.producer.Close()

	log.Printf("📤 Starting message publishing to topic: %s", kp.topic)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// Graceful shutdown support
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	i := 0
	for {
		select {
		case <-ticker.C:
			msg := &sarama.ProducerMessage{
				Topic: kp.topic,
				Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", i)),
				Value: sarama.StringEncoder(fmt.Sprintf("message #%d", i)),
			}
			partition, offset, err := kp.producer.SendMessage(msg)
			if err != nil {
				log.Printf("❌ Error sending message: %v", err)
				continue
			}
			log.Printf("✅ Sent message #%d to partition %d at offset %d", i, partition, offset)
			i++
		case <-signals:
			log.Println("🛑 Graceful shutdown triggered")
			return
		}
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
