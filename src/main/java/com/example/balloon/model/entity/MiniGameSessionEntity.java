package com.example.balloon.model.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.*;
import lombok.*;
import org.hibernate.annotations.CreationTimestamp;

@Entity
@Table(name = "mini_game_sessions")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class MiniGameSessionEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "id")
    private String id;

    @Column(name = "user_name")
    private String userName;

    @Column(name = "score")
    private Integer score;

    @CreationTimestamp
    @Column(name = "started_at", updatable = false)
    private String startedAt;

    @Column(name = "finished")
    private Boolean finished;
}