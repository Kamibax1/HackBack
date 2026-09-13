package com.example.balloon.repository;

import com.example.balloon.model.entity.RewardEntity;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface RewardRepository extends JpaRepository<RewardEntity, String> {
    List<RewardEntity> findByUserNameIgnoreCase(String username);
}