package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class FinishMiniGameRequest {
    @JsonProperty("session_id")
    private String sessionId;

    private Integer score;
}
