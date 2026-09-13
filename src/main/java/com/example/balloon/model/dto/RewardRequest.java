package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class RewardRequest {
    private String name;
    private Boolean claimed;

    @JsonProperty("user_name")
    private String userName;
}
