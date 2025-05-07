package com.library.management.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import com.library.management.entity.Notification;
import com.library.management.repo.NotificationRepository;

@Service
public class NotificationService {

    @Autowired
    private NotificationRepository repository;

    public void saveNotification(Notification notification) {
        repository.save(notification);
    }

    public List<Notification> getNotificationsByUser(Long userId) {
        return repository.findByUserId(userId);
    }

    public Page<Notification> getNotificationsByUser(Long userId, int page, int size) {
        PageRequest pageable = PageRequest.of(page, size, Sort.by("timestamp").descending());
        return repository.findByUserId(userId, pageable);
    }

    public void deleteNotification(Long id) {
        repository.deleteById(id);
    }

    public void markAsSeen(Long id) {
        repository.findById(id).ifPresent(notification -> {
            notification.setSeen(true);
            repository.save(notification);
        });
    }
}
