package com.kafka.kafka_consumer.configurations;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Component
@ConfigurationProperties(prefix = "spring.kafka.consumer")
public class KafkaSettingsService {
    private String topic;
    private String group;

    public void setTopic(String topic){
        this.topic= topic;
    }

    public void setGroup(String group){
        this.group= group;
    }

    public String getGroup() {
        return group;
    }

    public String getTopic() {
        return topic;
    }
}
