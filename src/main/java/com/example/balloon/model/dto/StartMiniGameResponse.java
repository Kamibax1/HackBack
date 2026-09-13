package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class StartMiniGameResponse {
    @JsonProperty("session_id")
    private String sessionId;

    private Integer duration;
}
