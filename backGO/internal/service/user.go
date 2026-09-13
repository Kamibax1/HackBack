package service

import (
	"Otbor/internal/database"
	"Otbor/internal/middleware"
	"Otbor/internal/models"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Авторизация пользователя
// @Description Вход по логину и паролю
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.UserLogin true "Login data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	var req models.UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	var user models.User

	if err := database.DB.
		Where("user_name = ?", req.UserName).
		First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
		return
	}

	salt, _ := base64.StdEncoding.DecodeString(user.Salt)

	if HashPassword(req.UserPassword, salt) != user.PasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong password",
		})
		return
	}

	token, err := middleware.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  user.UserRole,
	})
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Создание нового пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.UserRegister true "Register data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/register [post]
func Register(c *gin.Context) {
	var req models.UserRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	var existing models.User

	err := database.DB.
		Where("user_name = ?", req.UserName).
		First(&existing).Error

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "user already exists",
		})
		return
	}

	salt := make([]byte, 16)

	_, err = rand.Read(salt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot generate salt",
		})
		return
	}

	user := models.User{
		UserName:     req.UserName,
		UserEmail:    req.UserEmail,
		UserRole:     "User",
		PasswordHash: HashPassword(req.UserPassword, salt),
		Salt:         base64.StdEncoding.EncodeToString(salt),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered",
	})
}

// Update godoc
// @Summary Обновить пользователя
// @Description Изменение логина и пароля пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.UserUpdate true "Update data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/user/update [post]
func Update(c *gin.Context) {
	var req models.UserUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	var user models.User

	if err := database.DB.
		Where("user_name = ?", req.UserName).
		First(&user).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	salt, _ := base64.StdEncoding.DecodeString(user.Salt)

	if HashPassword(req.UserPassword, salt) != user.PasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong password",
		})
		return
	}

	if req.NewUserName != req.UserName {
		var existing models.User

		if err := database.DB.
			Where("user_name = ?", req.NewUserName).
			First(&existing).Error; err == nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "new username already exists",
			})
			return
		}
	}

	newSalt := make([]byte, 16)
	_, _ = rand.Read(newSalt)

	user.UserName = req.NewUserName
	user.PasswordHash = HashPassword(
		req.NewUserPassword,
		newSalt,
	)
	user.Salt = base64.StdEncoding.EncodeToString(newSalt)

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
}

// Delete godoc
// @Summary Удалить пользователя
// @Description Удаление пользователя по логину и паролю
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.UserDelete true "Delete data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/user/delete [delete]
func Delete(c *gin.Context) {
	var req models.UserDelete

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	var user models.User

	if err := database.DB.
		Where("user_name = ?", req.UserName).
		First(&user).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	salt, _ := base64.StdEncoding.DecodeString(user.Salt)

	if HashPassword(req.UserPassword, salt) != user.PasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong password",
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}

// Profile godoc
// @Summary Получить профиль пользователя
// @Description Получение профиля по имени пользователя
// @Tags Users
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.UserProfile
// @Failure 404 {object} map[string]string
// @Router /api/user/profile/{username} [get]
func Profile(c *gin.Context) {
	username := c.Param("username")

	var user models.User

	if err := database.DB.
		Where("user_name = ?", username).
		First(&user).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.UserProfile{
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
	})
}
