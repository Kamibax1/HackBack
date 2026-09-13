package com.example.balloon.repository;

import com.example.balloon.model.entity.TournamentEntity;
import org.springframework.data.jpa.repository.JpaRepository;

public interface TournamentRepository extends JpaRepository<TournamentEntity, String> {
}