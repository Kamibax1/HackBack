package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ActiveTournamentResponse {
    private String id;
    private String name;

    @JsonProperty("started_at")
    private LocalDateTime startedAt;

    @JsonProperty("ends_at")
    private LocalDateTime endsAt;
}
