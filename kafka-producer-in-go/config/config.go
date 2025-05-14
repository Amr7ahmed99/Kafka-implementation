// kafka/config.go
package kafka

import (
	"github.com/IBM/sarama"
)

// NewSaramaConfig creates a new Sarama configuration with sensible defaults holds the configuration for Kafka
func NewKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	// Set reasonable defaults for a reliable producer
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	return config
}
