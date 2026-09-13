package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class UserRegisterRequest {
    @JsonProperty("user_name")
    private String userName;

    @JsonProperty("user_password")
    private String userPassword;

    @JsonProperty("user_email")
    private String userEmail;
}
