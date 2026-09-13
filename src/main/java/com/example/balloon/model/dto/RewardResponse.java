package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class RewardResponse {
    private String id;
    private String name;
    private Boolean claimed;

    @JsonProperty("user_name")
    private String userName;
}
