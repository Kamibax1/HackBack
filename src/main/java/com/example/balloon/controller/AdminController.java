package com.example.balloon.controller;

import com.example.balloon.model.dto.*;
import com.example.balloon.service.AdminService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/admin")
public class AdminController {

    private final AdminService adminService;

    public AdminController(AdminService adminService) {
        this.adminService = adminService;
    }

    @GetMapping("/users")
    public ResponseEntity<List<UserResponse>> getUsers() {
        return ResponseEntity.ok(adminService.getAllUsers());
    }

    @GetMapping("/users/{username}")
    public ResponseEntity<UserResponse> getUser(@PathVariable String username) {
        return ResponseEntity.ok(adminService.getUser(username));
    }

    @PutMapping("/users/role")
    public ResponseEntity<UserResponse> changeRole(@RequestBody ChangeRoleRequest request) {
        return ResponseEntity.ok(adminService.changeRole(request));
    }

    @DeleteMapping("/users/{username}")
    public ResponseEntity<String> deleteUser(@PathVariable String username) {
        return ResponseEntity.ok(adminService.deleteUser(username));
    }

    @GetMapping("/tournaments")
    public ResponseEntity<List<ActiveTournamentResponse>> getTournaments() {
        return ResponseEntity.ok(adminService.getTournaments());
    }

    @PostMapping("/tournaments")
    public ResponseEntity<ActiveTournamentResponse> createTournament(@RequestBody CreateTournamentRequest request) {
        return ResponseEntity.status(HttpStatus.CREATED).body(adminService.createTournament(request));
    }

    @DeleteMapping("/tournaments/{id}")
    public ResponseEntity<String> deleteTournament(@PathVariable String id) {
        return ResponseEntity.ok(adminService.deleteTournament(id));
    }

    @PostMapping("/rewards")
    public ResponseEntity<RewardResponse> createReward(@RequestBody RewardRequest reward) {
        return ResponseEntity.status(HttpStatus.CREATED).body(adminService.createReward(reward));
    }

    @GetMapping("/rewards")
    public ResponseEntity<List<RewardResponse>> getRewards() {
        return ResponseEntity.ok(adminService.getRewards());
    }

    @GetMapping("/games")
    public ResponseEntity<List<GameHistoryResponse>> getGames() {
        return ResponseEntity.ok(adminService.getGames());
    }
}