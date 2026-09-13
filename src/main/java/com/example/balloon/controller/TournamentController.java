package com.example.balloon.controller;

import com.example.balloon.model.dto.ActiveTournamentResponse;
import com.example.balloon.model.dto.LeaderboardEntryResponse;
import com.example.balloon.model.dto.MiniGameLeaderboardProjection;
import com.example.balloon.model.dto.TournamentResponse;
import com.example.balloon.service.TournamentService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api")
public class TournamentController {

    private final TournamentService tournamentService;

    public TournamentController(TournamentService tournamentService) {
        this.tournamentService = tournamentService;
    }

    @GetMapping("/tournament")
    public ResponseEntity<List<ActiveTournamentResponse>> getTournaments() {
        return ResponseEntity.ok(tournamentService.getActiveTournaments());
    }

    @GetMapping("/latest_tournaments")
    public ResponseEntity<List<TournamentResponse>> getLatestTournaments() {
        return ResponseEntity.ok(tournamentService.getLatestTournaments());
    }

    @GetMapping("/tournament/top")
    public ResponseEntity<List<MiniGameLeaderboardProjection>> getLeaderBoard() {
        return ResponseEntity.ok(tournamentService.getLeaderboard());
    }
}
