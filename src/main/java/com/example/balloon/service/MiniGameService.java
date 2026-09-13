package com.example.balloon.service;

import com.example.balloon.exception.BadRequestException;
import com.example.balloon.exception.NotFoundException;
import com.example.balloon.model.dto.*;
import com.example.balloon.model.entity.GameHistoryEntity;
import com.example.balloon.model.entity.MiniGameSessionEntity;
import com.example.balloon.repository.GameHistoryRepository;
import com.example.balloon.repository.MiniGameSessionRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.ZonedDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;

@Service
public class MiniGameService {

    private final MiniGameSessionRepository sessionRepository;
    private final GameHistoryRepository gameHistoryRepository;

    public MiniGameService(MiniGameSessionRepository sessionRepository,
                           GameHistoryRepository gameHistoryRepository) {
        this.sessionRepository = sessionRepository;
        this.gameHistoryRepository = gameHistoryRepository;
    }

    @Transactional
    public StartMiniGameResponse start(String userName) {
        MiniGameSessionEntity session = new MiniGameSessionEntity();
        session.setUserName(userName);
        session.setScore(0);
        session.setStartedAt(ZonedDateTime.now().format(DateTimeFormatter.ISO_INSTANT));
        session.setFinished(false);

        sessionRepository.save(session);

        return new StartMiniGameResponse(session.getId(), 30);
    }

    @Transactional
    public FinishMiniGameResponse finish(FinishMiniGameRequest req) {
        MiniGameSessionEntity session = sessionRepository.findById(req.getSessionId())
                .orElseThrow(() -> new NotFoundException("session not found"));

        if (session.getFinished()) {
            throw new BadRequestException("session already finished");
        }

        session.setScore(req.getScore());
        session.setFinished(true);
        sessionRepository.save(session);

        GameHistoryEntity history = new GameHistoryEntity();
        history.setUserName(session.getUserName());
        history.setScore(req.getScore());
        history.setPlayedAt(ZonedDateTime.now().format(DateTimeFormatter.ISO_INSTANT));
        gameHistoryRepository.save(history);

        FinishMiniGameResponse finishMiniGameResponse = new FinishMiniGameResponse();
        finishMiniGameResponse.setSessionId(req.getSessionId());
        finishMiniGameResponse.setScore(req.getScore());
        return finishMiniGameResponse;
    }

    @Transactional(readOnly = true)
    public List<MiniGameLeaderboardProjection> getLeaderboard() {
        return sessionRepository.findLeaderboard();
    }
}