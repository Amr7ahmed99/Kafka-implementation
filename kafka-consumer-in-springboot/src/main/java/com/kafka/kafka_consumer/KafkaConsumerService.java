package com.kafka.kafka_consumer;

import com.kafka.kafka_consumer.configurations.KafkaSettingsService;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Random;

@Service
public class KafkaConsumerService {

    private static final Logger logger = LoggerFactory.getLogger(KafkaConsumerService.class);
    private final Random random = new Random();
    private final KafkaSettingsService kafkaSettingsService;

    public KafkaConsumerService(KafkaSettingsService kafkaSettingsService){
        this.kafkaSettingsService= kafkaSettingsService;
    }


    @KafkaListener(topics = "#{@kafkaSettingsService.getTopic()}", groupId = "#{@kafkaSettingsService.getGroup()}")
    public void consume(ConsumerRecord<String, String> record) {
        try {
            String timestamp = new SimpleDateFormat("HH:mm:ss.SSS").format(new Date());
            logger.info("📥 [Group: {}] Consumed '{}' from partition {} @ offset {} at {}",
                    kafkaSettingsService.getGroup(), record.value(), record.partition(), record.offset(), timestamp);
            Thread.sleep(random.nextInt(100));
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            logger.error("⚠️ Interrupted while processing message: {}", e.getMessage());
        }
    }
}
