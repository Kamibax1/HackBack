package com.example.balloon.config;

import com.example.balloon.model.entity.UserEntity;
import com.example.balloon.model.enums.RoleEnum;
import com.example.balloon.repository.UserRepository;
import com.example.balloon.service.PasswordService;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
public class DataSeeder implements CommandLineRunner {

    private final UserRepository userRepository;
    private final PasswordService passwordService;

    public DataSeeder(UserRepository userRepository, PasswordService passwordService) {
        this.userRepository = userRepository;
        this.passwordService = passwordService;
    }

    @Override
    public void run(String... args) {
        if (userRepository.findByUserName("admin").isEmpty()) {
            byte[] salt = passwordService.generateSalt();

            UserEntity admin = new UserEntity();
            admin.setUserName("admin");
            admin.setUserEmail("admin@localhost");
            admin.setUserRole(RoleEnum.ADMIN);
            admin.setPasswordHash(passwordService.hashPassword("admin123", salt));
            admin.setSalt(passwordService.encodeSalt(salt));

            userRepository.save(admin);
        }
    }
}