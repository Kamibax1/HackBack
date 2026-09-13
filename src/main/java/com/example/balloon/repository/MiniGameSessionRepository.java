package com.example.balloon.repository;

import com.example.balloon.model.dto.MiniGameLeaderboardProjection;
import com.example.balloon.model.entity.MiniGameSessionEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface MiniGameSessionRepository extends JpaRepository<MiniGameSessionEntity, String> {

    @Query(value = """
            SELECT 
                user_name AS userName, 
                MAX(score) AS score 
            FROM mini_game_sessions 
            WHERE finished = true 
            GROUP BY user_name 
            ORDER BY score DESC 
            LIMIT 100
            """,
            nativeQuery = true)
    List<MiniGameLeaderboardProjection> findLeaderboard();
}