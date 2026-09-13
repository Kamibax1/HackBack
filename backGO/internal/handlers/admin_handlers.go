package handlers

import (
	"Otbor/internal/database"

	"Otbor/internal/models"
	"Otbor/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ChangeRoleRequest struct {
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

type CreateTournamentRequest struct {
	Name      string    `json:"name"`
	StartedAt time.Time `json:"startedAt"`
	EndsAt    time.Time `json:"endedAt"`
}

func parseTournamentTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)

	if err != nil {
		return time.Now().Add(24 * time.Hour)
	}

	return t
}

// AdminGetUsers godoc
// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.UserResponse
// @Failure 500 {object} map[string]string
// @Router /api/admin/users [get]
func AdminGetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := make([]models.UserResponse, 0, len(users))

	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:        u.ID,
			UserName:  u.UserName,
			UserEmail: u.UserEmail,
			UserRole:  u.UserRole,
		})
	}

	c.JSON(http.StatusOK, response)
}

// AdminGetUser godoc
// @Summary Получить пользователя
// @Description Получить пользователя по username
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param username path string true "Username"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /api/admin/users/{username} [get]
func AdminGetUser(c *gin.Context) {
	var user models.User

	if err := database.DB.
		Where("user_name = ?", c.Param("username")).
		First(&user).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// AdminChangeRole godoc
// @Summary Изменить роль пользователя
// @Description Изменение роли пользователя
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body ChangeRoleRequest true "Role data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/admin/users/role [put]
func AdminChangeRole(c *gin.Context) {
	var req ChangeRoleRequest

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

	user.UserRole = req.Role

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// AdminDeleteUser godoc
// @Summary Удалить пользователя
// @Description Удаление пользователя по username
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param username path string true "Username"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/admin/users/{username} [delete]
func AdminDeleteUser(c *gin.Context) {
	var user models.User

	if err := database.DB.
		Where("user_name = ?", c.Param("username")).
		First(&user).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
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

// AdminCreateTournament godoc
// @Summary Создать турнир
// @Description Создание нового турнира
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateTournamentRequest true "Tournament"
// @Success 201 {object} CreateTournamentRequest
// @Failure 400 {object} map[string]string
// @Router /api/admin/tournaments [post]
func AdminCreateTournament(c *gin.Context) {
	var req CreateTournamentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tournament := service.ActiveTournament{
		ID:        req.Name,
		Name:      req.Name,
		StartedAt: req.StartedAt,
		EndsAt:    req.EndsAt,
	}

	if err := service.SetActiveTournament(tournament); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, tournament)
}

// AdminGetTournaments godoc
// @Summary Список турниров
// @Description Получить все турниры
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} CreateTournamentRequest
// @Router /api/latest_tournaments [get]
func GetLatestTournaments(c *gin.Context) {
	var tournaments []models.Tournament

	if err := database.DB.Find(&tournaments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

// AdminGetTournaments godoc
// @Summary Список турниров
// @Description Получить все турниры
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} CreateTournamentRequest
// @Router /api/latest_tournaments [get]
func GetTournaments(c *gin.Context) {
	tournament, err := service.GetActiveTournament()

	if err != nil {
		c.JSON(http.StatusOK, []service.ActiveTournament{})
		return
	}

	c.JSON(
		http.StatusOK,
		[]service.ActiveTournament{
			*tournament,
		},
	)
}

// AdminDeleteTournament godoc
// @Summary Удалить турнир
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "Tournament ID"
// @Success 200 {object} map[string]string
// @Router /api/admin/tournaments/{id} [delete]
func AdminDeleteTournament(c *gin.Context) {
	if err := database.DB.Delete(
		&models.Tournament{},
		c.Param("id"),
	).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tournament deleted",
	})
}

// AdminCreateReward godoc
// @Summary Создать награду
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.Reward true "Reward"
// @Success 201 {object} models.Reward
// @Router /api/admin/rewards [post]
func AdminCreateReward(c *gin.Context) {
	var reward models.Reward

	if err := c.ShouldBindJSON(&reward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if err := database.DB.Create(&reward).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, reward)
}

// AdminGetRewards godoc
// @Summary Все награды
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Reward
// @Router /api/admin/rewards [get]
func AdminGetRewards(c *gin.Context) {
	var rewards []models.Reward

	if err := database.DB.Find(&rewards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rewards)
}

// AdminGetGames godoc
// @Summary История игр
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.GameHistory
// @Router /api/admin/games [get]
func AdminGetGames(c *gin.Context) {
	var games []models.GameHistory

	if err := database.DB.
		Order("played_at DESC").
		Find(&games).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, games)
}
