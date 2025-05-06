package com.library.management.entity;

import java.time.LocalDateTime;
import java.util.List;

import lombok.Data;

@Data
public class ApiError {
	
    private LocalDateTime timestamp;
    private int status;
    private String error;
    private List<String> messages;

    public ApiError(int status, String error, List<String> messages) {
        this.timestamp = LocalDateTime.now();
        this.status = status;
        this.error = error;
        this.messages = messages;
    }

}
