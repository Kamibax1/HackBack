package com.example.balloon.service;

import com.example.balloon.exception.NotFoundException;
import com.example.balloon.model.dto.*;
import com.example.balloon.model.entity.RewardEntity;
import com.example.balloon.model.entity.TournamentEntity;
import com.example.balloon.model.entity.UserEntity;
import com.example.balloon.model.mapper.EntityMapper;
import com.example.balloon.repository.GameHistoryRepository;
import com.example.balloon.repository.RewardRepository;
import com.example.balloon.repository.TournamentRepository;
import com.example.balloon.repository.UserRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class AdminService {
    private final UserRepository userRepository;
    private final TournamentRepository tournamentRepository;
    private final RewardRepository rewardRepository;
    private final GameHistoryRepository gameHistoryRepository;
    private final EntityMapper mapper;

    public AdminService(UserRepository userRepository, TournamentRepository tournamentRepository, RewardRepository rewardRepository, GameHistoryRepository gameHistoryRepository, EntityMapper mapper) {
        this.userRepository = userRepository;
        this.tournamentRepository = tournamentRepository;
        this.rewardRepository = rewardRepository;
        this.gameHistoryRepository = gameHistoryRepository;
        this.mapper = mapper;
    }

    @Transactional(readOnly = true)
    public List<UserResponse> getAllUsers() {
        return mapper.toUserResponseList(userRepository.findAll());
    }

    @Transactional(readOnly = true)
    public UserResponse getUser(String userName) {
        return userRepository.findByUserName(userName)
                .map(mapper::toUserResponse)
                .orElseThrow(() -> new NotFoundException("user not found"));
    }

    @Transactional
    public UserResponse changeRole(ChangeRoleRequest req) {
        UserEntity user = userRepository.findByUserName(req.getUserName())
                .orElseThrow(() -> new NotFoundException("user not found"));
        user.setUserRole(req.getRole());
        return mapper.toUserResponse(userRepository.save(user));
    }

    @Transactional
    public String deleteUser(String userName) {
        UserEntity user = userRepository.findByUserName(userName)
                .orElseThrow(() -> new NotFoundException("user not found"));
        userRepository.delete(user);
        return "user deleted";
    }

    @Transactional(readOnly = true)
    public List<ActiveTournamentResponse> getTournaments() {
        return mapper.toActiveTournamentResponseList(tournamentRepository.findAll());
    }

    @Transactional
    public ActiveTournamentResponse createTournament(CreateTournamentRequest request) {
        TournamentEntity newTournament = new TournamentEntity();
        newTournament.setName(request.getName());
        newTournament.setStartedAt(request.getStartedAt());
        newTournament.setEndedAt(request.getEndsAt());
        TournamentEntity saved = tournamentRepository.save(newTournament);
        return mapper.toActiveTournamentResponse(saved);
    }

    @Transactional
    public String deleteTournament(String id) {
        TournamentEntity tournamentEntity = tournamentRepository.findById(id)
                .orElseThrow(() -> new NotFoundException("tournament not found"));
        tournamentRepository.delete(tournamentEntity);
        return "tournament deleted";
    }

    @Transactional
    public RewardResponse createReward(RewardRequest request) {
        RewardEntity newReward = new RewardEntity();
        newReward.setName(request.getName());
        newReward.setClaimed(request.getClaimed());
        newReward.setUserName(request.getUserName());
        return mapper.toRewardResponse(rewardRepository.save(newReward));
    }

    @Transactional
    public List<RewardResponse> getRewards() {
        return mapper.toRewardResponseList(rewardRepository.findAll());
    }

    @Transactional
    public List<GameHistoryResponse> getGames() {
        return mapper.toGameHistoryResponseList(gameHistoryRepository.findAll());
    }
}