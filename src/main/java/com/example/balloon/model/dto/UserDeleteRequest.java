package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class UserDeleteRequest {
    @JsonProperty("user_name")
    private String userName;

    @JsonProperty("user_password")
    private String userPassword;
}
