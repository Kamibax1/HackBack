package com.example.balloon.model.dto;

import lombok.Data;
import java.time.LocalDateTime;

@Data
public class CreateTournamentRequest {
    private String name;
    private LocalDateTime startedAt;
    private LocalDateTime endsAt;
}
