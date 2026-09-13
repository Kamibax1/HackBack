package com.example.balloon.repository;

import com.example.balloon.model.entity.UserEntity;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface UserRepository extends JpaRepository<UserEntity, String> {
    Optional<UserEntity> findByUserName(String userName);
    boolean existsByUserName(String userName);
}