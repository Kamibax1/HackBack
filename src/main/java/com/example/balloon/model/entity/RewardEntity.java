package com.example.balloon.model.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name = "rewards")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class RewardEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @JsonProperty("id")
    private String id;

    @Column(name = "name")
    private String name;

    @Column(name = "claimed")
    private Boolean claimed;

    @Column(name = "user_name")
    private String userName;
}