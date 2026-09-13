package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public interface MiniGameLeaderboardProjection {

    @JsonProperty("user_name")
    String getUserName();

    Integer getScore();
}