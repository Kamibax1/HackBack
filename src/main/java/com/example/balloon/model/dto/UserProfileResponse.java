package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class UserProfileResponse {
    @JsonProperty("user_name")
    private String userName;

    @JsonProperty("user_email")
    private String userEmail;
}
