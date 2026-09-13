package com.example.balloon.controller;

import com.example.balloon.model.dto.*;
import com.example.balloon.service.MiniGameService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.security.Principal;
import java.util.List;

@RestController
@RequestMapping("/api/minigame")
public class MiniGameController {

    private final MiniGameService miniGameService;

    public MiniGameController(MiniGameService miniGameService) {
        this.miniGameService = miniGameService;
    }

    @PostMapping("/start")
    public ResponseEntity<StartMiniGameResponse> start(Principal principal) {
        if (principal == null || principal.getName() == null) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).build();
        }

        StartMiniGameResponse response = miniGameService.start(principal.getName());
        return ResponseEntity.ok(response);
    }

    @PostMapping("/finish")
    public ResponseEntity<FinishMiniGameResponse> finish(@RequestBody FinishMiniGameRequest request) {
        return ResponseEntity.ok(miniGameService.finish(request));
    }

    @GetMapping("/leaderboard")
    public ResponseEntity<List<MiniGameLeaderboardProjection>> getLeaderboard() {
        return ResponseEntity.ok(miniGameService.getLeaderboard());
    }
}