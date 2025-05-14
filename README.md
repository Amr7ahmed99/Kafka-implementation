# Apache Kafka Integration with Go and Spring Boot

This project demonstrates a simple Kafka-based messaging system with two microservices:

- **Producer Service** written in **Go**
- **Consumer Service** written in **Spring Boot (Java)**

The services communicate through an Apache Kafka topic to simulate message publishing and consumption, Messages are sent from the Go service to an Apache Kafka topic, and consumed dynamically by the Spring Boot service.

---

## 🛠️ Technologies Used

- [Apache Kafka](https://kafka.apache.org/)
- Go (with [sarama](https://github.com/IBM/sarama) Kafka client)
- Java with [Spring Boot](https://spring.io/projects/spring-boot)
- [Spring Kafka](https://spring.io/projects/spring-kafka)

---

## 📌 Architecture Overview
+-----------------+ Kafka Broker +-------------------------+
| Go Producer | ---> my-kafka-topic ---> | Spring Consumer |
+-----------------+ +------------------+-------------------+

---

## 📂 Project Structure

### Go Producer (`/kafka`)

- `producer.go`: Sets up and runs the Kafka message producer.
- `config.go`: Provides the Kafka configuration using the Sarama library.

#### Environment variables:
| Variable       | Default         | Description                         |
|----------------|-----------------|-------------------------------------|
| `KAFKA_HOST`   | `localhost:9092`| Kafka broker address                |
| `KAFKA_TOPIC`  | `my-kafka-topic`   | Kafka topic to publish messages to |


### Spring Boot Consumer

- `KafkaConsumerService.java`: Consumes messages dynamically from Kafka.
- `KafkaSettingsService.java`: Injects Kafka topic and group ID from properties.

#### The following properties are used to configure the Spring Boot consumer:

```bash
    # Kafka bootstrap server
    spring.kafka.bootstrap-servers=localhost:9092

    # Kafka consumer group id
    spring.kafka.consumer.group=my-kafka-group

    # Kafka consumer topic id
    spring.kafka.consumer.topic=my-kafka-topic

    # Kafka deserializer classes
    spring.kafka.consumer.key-deserializer=org.apache.kafka.common.serialization.StringDeserializer
    spring.kafka.consumer.value-deserializer=org.apache.kafka.common.serialization.StringDeserializer

    # Start consuming from the beginning if no offset is found
    spring.kafka.consumer.auto-offset-reset=earliest
```

---

## 🚀 How to Run

### Prerequisites
. Kafka and Zookeeper running locally or on a reachable server.

. Go installed (v1.18+ recommended)

. Java 17+ and Maven

1. Start Kafka and Zookeeper
You can use Docker Or install locally from the - [Apache Kafka website](https://kafka.apache.org/quickstart):
```bash
    docker-compose up -d
```

2. Create kafka topic
```bash
    kafka-topics.sh --create --topic :topic-name --partitions :#partitions --replication-factor :#replications --zookeeper zookeeper:2181
```

3. Run the Go Producer
It will start sending messages every second to my-kafka-topic.
```bash
    cd path/to/your/go/project
    go run main.go
```

4. Run the Spring Boot Consumer
It will listen on the same topic and print out consumed messages.
```bash
    cd path/to/your/springboot/project
    ./mvnw spring-boot:run
```

---

## 📝 Sample Output

### Producer:
``` bash
    📤 Starting message publishing to topic: my-kafka-topic
    ✅ Sent message #0 to partition 0 at offset 1
```

### Consumer:
``` bash
    📥 [Group: my-kafka-group] Consumed 'message #0' from partition 0 @ offset 1 at 12:45:03.456
```

---