package com.example.balloon.model.mapper;

import com.example.balloon.model.dto.*;
import com.example.balloon.model.entity.*;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class EntityMapper {
    public TournamentResponse toTournamentResponse(TournamentEntity entity) {
        return new TournamentResponse(
                entity.getId(),
                entity.getName(),
                entity.getStartedAt(),
                entity.getEndedAt(),
                entity.getWinner()
        );
    }

    public List<TournamentResponse> toTournamentResponse(List<TournamentEntity> entities) {
        return entities.stream().map(this::toTournamentResponse).toList();
    }

    public ActiveTournamentResponse toActiveTournamentResponse(TournamentEntity entity) {
        return new ActiveTournamentResponse(
                entity.getId(),
                entity.getName(),
                entity.getStartedAt(),
                entity.getEndedAt()
        );
    }

    public List<ActiveTournamentResponse> toActiveTournamentResponseList(List<TournamentEntity> entities) {
        return entities.stream().map(this::toActiveTournamentResponse).toList();
    }

    public RewardResponse toRewardResponse(RewardEntity entity) {
        return new RewardResponse(
                entity.getId(),
                entity.getName(),
                entity.getClaimed(),
                entity.getUserName()
        );
    }

    public List<RewardResponse> toRewardResponseList(List<RewardEntity> entities) {
        return entities.stream().map(this::toRewardResponse).toList();
    }

    public UserResponse toUserResponse(UserEntity entity) {
        return new UserResponse(
                entity.getId(),
                entity.getUserName(),
                entity.getUserEmail(),
                entity.getUserRole()
        );
    }

    public List<UserResponse> toUserResponseList(List<UserEntity> entities) {
        return entities.stream().map(this::toUserResponse).toList();
    }

    public GameHistoryResponse toGameHistoryResponse(GameHistoryEntity entity) {
        return new GameHistoryResponse(
                entity.getId(),
                entity.getUserName(),
                entity.getScore(),
                entity.getPlayedAt()
        );
    }

    public List<GameHistoryResponse> toGameHistoryResponseList(List<GameHistoryEntity> entities) {
        return entities.stream().map(this::toGameHistoryResponse).toList();
    }
}