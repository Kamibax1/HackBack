package com.example.balloon.service;

import com.example.balloon.exception.*;
import com.example.balloon.model.dto.*;
import com.example.balloon.model.entity.UserEntity;
import com.example.balloon.model.enums.RoleEnum;
import com.example.balloon.repository.UserRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class UserService {

    private final UserRepository userRepository;
    private final PasswordService passwordService;
    private final JwtService jwtService;

    public UserService(UserRepository userRepository, PasswordService passwordService, JwtService jwtService) {
        this.userRepository = userRepository;
        this.passwordService = passwordService;
        this.jwtService = jwtService;
    }

    @Transactional(readOnly = true)
    public UserLoginResponse login(UserLoginRequest request) {
        UserEntity user = userRepository.findByUserName(request.getUserName())
                .orElseThrow(() -> new UnauthorizedException("user not found"));

        if (!checkPassword(request.getUserPassword(), user)) {
            throw new UnauthorizedException("wrong password");
        }

        String token = jwtService.generateToken(user);
        return new UserLoginResponse(token, user.getUserRole());
    }

    @Transactional
    public String register(UserRegisterRequest request) {
        if (userRepository.existsByUserName(request.getUserName())) {
            throw new ConflictException("user already exists");
        }
        byte[] salt = passwordService.generateSalt();

        UserEntity user = new UserEntity();
        user.setUserName(request.getUserName());
        user.setUserEmail(request.getUserEmail());
        user.setPasswordHash(passwordService.hashPassword(request.getUserPassword(), salt));
        user.setSalt(passwordService.encodeSalt(salt));
        user.setUserRole(RoleEnum.USER);

        userRepository.save(user);
        return "user registered";
    }

    @Transactional
    public String update(UserUpdateRequest request) {
        UserEntity user = userRepository.findByUserName(request.getUserName())
                .orElseThrow(() -> new NotFoundException("user not found"));

        if (!checkPassword(request.getUserPassword(), user)) {
            throw new UnauthorizedException("wrong password");
        }

        String newUserName = request.getNewUserName() != null ? request.getNewUserName() : user.getUserName();
        if (!newUserName.equals(user.getUserName()) && userRepository.existsByUserName(newUserName)) {
            throw new BadRequestException("new username already exists");
        }
        user.setUserName(newUserName);

        if (request.getNewUserPassword() != null && !request.getNewUserPassword().isBlank()) {
            byte[] newSalt = passwordService.generateSalt();
            user.setPasswordHash(passwordService.hashPassword(request.getNewUserPassword(), newSalt));
            user.setSalt(passwordService.encodeSalt(newSalt));
        }

        userRepository.save(user);
        return "user updated";
    }

    @Transactional
    public String delete(UserDeleteRequest request) {
        UserEntity user = userRepository.findByUserName(request.getUserName())
                .orElseThrow(() -> new NotFoundException("user not found"));

        if (!checkPassword(request.getUserPassword(), user)) {
            throw new UnauthorizedException("wrong password");
        }
        userRepository.delete(user);
        return "user deleted";
    }

    @Transactional(readOnly = true)
    public UserProfileResponse getProfile(String userName) {
        UserEntity user = userRepository.findByUserName(userName)
                .orElseThrow(() -> new NotFoundException("user not found"));
        return new UserProfileResponse(user.getUserName(), user.getUserEmail());
    }

    @Transactional(readOnly = true)
    public java.util.List<UserEntity> getAllUsers() {
        return userRepository.findAll();
    }

    @Transactional(readOnly = true)
    public UserEntity getUserByUserName(String userName) {
        return userRepository.findByUserName(userName)
                .orElseThrow(() -> new NotFoundException("user not found"));
    }

    @Transactional
    public void changeRole(String userName, RoleEnum newRole) {
        UserEntity user = getUserByUserName(userName);
        user.setUserRole(newRole);
        userRepository.save(user);
    }

    private boolean checkPassword(String rawPassword, UserEntity user) {
        byte[] salt = passwordService.decodeSalt(user.getSalt());
        String hash = passwordService.hashPassword(rawPassword, salt);
        return java.security.MessageDigest.isEqual(
                hash.getBytes(java.nio.charset.StandardCharsets.UTF_8),
                user.getPasswordHash().getBytes(java.nio.charset.StandardCharsets.UTF_8)
        );
    }
}