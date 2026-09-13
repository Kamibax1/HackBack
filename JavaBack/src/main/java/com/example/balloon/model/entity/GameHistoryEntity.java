package com.example.balloon.model.entity;

import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name = "game_histories")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class GameHistoryEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "id")
    private String id;

    @Column(name = "user_name")
    private String userName;

    @Column(name = "score")
    private Integer score;

    @Column(name = "played_at")
    private String playedAt;
}