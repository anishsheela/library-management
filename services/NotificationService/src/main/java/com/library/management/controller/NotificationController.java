package com.library.management.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.library.management.entity.Notification;
import com.library.management.service.NotificationService;

import jakarta.validation.Valid;

@RestController
@RequestMapping("/notifications")
public class NotificationController {

    @Autowired
    private NotificationService service;

    @GetMapping("/{userId}")
    public Page<Notification> getNotifications(
            @PathVariable Long userId,
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "10") int size
    ) {
        return service.getNotificationsByUser(userId, page, size);
    }

    @PostMapping
    public ResponseEntity<Notification> createNotification(@Valid @RequestBody Notification notification) {
        service.saveNotification(notification);
        return ResponseEntity.ok(notification);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<?> deleteNotification(@PathVariable Long id) {
        service.deleteNotification(id);
        return ResponseEntity.ok().build();
    }

    @PatchMapping("/{id}/seen")
    public ResponseEntity<?> markAsSeen(@PathVariable Long id) {
        service.markAsSeen(id);
        return ResponseEntity.ok().build();
    }
}
