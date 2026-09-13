package com.example.balloon.service;

import com.example.balloon.model.dto.ActiveTournamentResponse;
import com.example.balloon.model.dto.LeaderboardEntryResponse;
import com.example.balloon.model.dto.MiniGameLeaderboardProjection;
import com.example.balloon.model.dto.TournamentResponse;
import com.example.balloon.model.mapper.EntityMapper;
import com.example.balloon.repository.GameHistoryRepository;
import com.example.balloon.repository.MiniGameSessionRepository;
import com.example.balloon.repository.TournamentRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class TournamentService {
    private final TournamentRepository tournamentRepository;
    private final GameHistoryRepository gameHistoryRepository;
    private final MiniGameSessionRepository miniGameSessionRepository;
    private final EntityMapper mapper;

    public TournamentService(TournamentRepository tournamentRepository, GameHistoryRepository gameHistoryRepository, MiniGameSessionRepository miniGameSessionRepository, EntityMapper mapper) {
        this.tournamentRepository = tournamentRepository;
        this.gameHistoryRepository = gameHistoryRepository;
        this.miniGameSessionRepository = miniGameSessionRepository;
        this.mapper = mapper;
    }

    @Transactional(readOnly = true)
    public List<ActiveTournamentResponse> getActiveTournaments() {
        return mapper.toActiveTournamentResponseList(tournamentRepository.findAll());
    }

    @Transactional(readOnly = true)
    public List<TournamentResponse> getLatestTournaments() {
        return mapper.toTournamentResponse(tournamentRepository.findAll());
    }

    @Transactional(readOnly = true)
    public List<MiniGameLeaderboardProjection> getLeaderboard() {
        return miniGameSessionRepository.findLeaderboard();
    }

    @Transactional
    public void delete(String id) {
        tournamentRepository.deleteById(id);
    }
}