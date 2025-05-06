package com.library.management.config;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.library.management.entity.Notification;
import com.library.management.service.NotificationService;

@Component
public class NotificationConsumer {

    @Autowired
    private NotificationService service;

    @KafkaListener(topics = {"book.due_date_reminder", "book.overdue", "book.available_reserved"}, groupId = "notification-group")
    public void listen(ConsumerRecord<String, String> record) {
        try {
            String value = record.value();
            ObjectMapper mapper = new ObjectMapper();
            JsonNode node = mapper.readTree(value);

            Notification notification = new Notification();
            notification.setUserId(node.get("userId").asLong());
            notification.setMessage(node.get("message").asText());
            service.saveNotification(notification);

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
