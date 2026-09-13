package com.example.balloon.controller;

import com.example.balloon.model.dto.*;
import com.example.balloon.service.UserService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/user")
public class UserController {

    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @PostMapping("/login")
    public ResponseEntity<UserLoginResponse> login(@RequestBody UserLoginRequest request) {
        return ResponseEntity.ok(userService.login(request));
    }

    @PostMapping("/register")
    public ResponseEntity<String> register(@RequestBody UserRegisterRequest request) {
        return ResponseEntity.status(HttpStatus.CREATED).body(userService.register(request));
    }

    @PostMapping("/update")
    public ResponseEntity<String> update(@RequestBody UserUpdateRequest request) {
        return ResponseEntity.ok(userService.update(request));
    }

    @DeleteMapping("/delete")
    public ResponseEntity<String> delete(@RequestBody UserDeleteRequest request) {
        return ResponseEntity.ok(userService.delete(request));
    }

    @GetMapping("/profile/{username}")
    public ResponseEntity<UserProfileResponse> profile(@PathVariable String username) {
        return ResponseEntity.ok(userService.getProfile(username));
    }
}
