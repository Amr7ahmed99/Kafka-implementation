# Configuration
TOPIC_NAME=my-kafka-topic
PARTITIONS=3
REPLICATION_FACTOR=1
ZOOKEEPER=zookeeper:2181
GO_PROJECT_PATH=./kafka-producer-in-go
SPRING_PROJECT_PATH=./kafka-consumer-in-springboot
KAFKA_CONTAINER_NAME=kafka


# Start Kafka and Zookeeper via Docker Compose
start-kafka:
	docker-compose up -d

# Open Kafka interactive terminal (interactive)
kafka-shell:
	docker exec -it $(KAFKA_CONTAINER_NAME) /bin/bash

# Create Kafka topic (runs inside Kafka container)
create-topic:
	docker exec -it $(KAFKA_CONTAINER_NAME) kafka-topics.sh \
		--create \
		--topic $(TOPIC_NAME) \
		--partitions $(PARTITIONS) \
		--replication-factor $(REPLICATION_FACTOR) \
		--zookeeper $(ZOOKEEPER)

# Run Go Producer
run-producer:
	cd $(GO_PROJECT_PATH) && go run main.go

# Run Spring Boot Consumer
run-consumer:
	cd $(SPRING_PROJECT_PATH) && ./mvnw spring-boot:run

# Run both producer and consumer
run-both:
	$(MAKE) run-producer & $(MAKE) run-consumer

start:
	$(MAKE) start-kafka
	$(MAKE) create-topic
	$(MAKE) run-both


.PHONY: start-kafka kafka-shell create-topic run-producer run-consumer run-both start
