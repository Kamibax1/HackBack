package com.example.balloon.model.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class UserUpdateRequest {
    @JsonProperty("user_name")
    private String userName;

    @JsonProperty("user_password")
    private String userPassword;

    @JsonProperty("new_user_name")
    private String newUserName;

    @JsonProperty("new_user_password")
    private String newUserPassword;
}
