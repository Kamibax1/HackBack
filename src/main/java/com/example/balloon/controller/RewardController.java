package com.example.balloon.controller;

import com.example.balloon.model.dto.ClaimRewardRequest;
import com.example.balloon.model.dto.RewardResponse;
import com.example.balloon.service.RewardService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/rewards")
public class RewardController {

    private final RewardService rewardService;

    public RewardController(RewardService rewardService) {
        this.rewardService = rewardService;
    }

    @GetMapping
    public ResponseEntity<List<RewardResponse>> getRewards(@RequestParam String username) {
        return ResponseEntity.ok(rewardService.getRewards(username));
    }

    @PostMapping("/claim")
    public ResponseEntity<RewardResponse> claimReward(@RequestBody ClaimRewardRequest request) {
        return ResponseEntity.ok(rewardService.claimReward(request));
    }
}
