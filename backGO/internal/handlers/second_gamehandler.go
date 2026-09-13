package handlers

import (
	"Otbor/internal/database"
	"Otbor/internal/models"
	"Otbor/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FinishMiniGameRequest struct {
	SessionID string `json:"session_id"`
	Score     int    `json:"score"`
	Hash      string `json:"hash"`
}

type MiniGameLeaderboardEntry struct {
	UserName string `json:"user_name"`
	Score    int    `json:"score"`
}

// StartMiniGame godoc
// @Summary Начать мини-игру
// @Tags MiniGame
// @Produce json
// @Security BearerAuth
// @Success 200 {object} StartMiniGameResponse
// @Router /api/minigame/start [post]
func StartMiniGame(c *gin.Context) {
	username := c.GetString("username")

	session := models.ActiveGameSession{
		ID:         uuid.NewString(),
		UserName:   username,
		StartedAt:  time.Now().Unix(),
		ServerSeed: uuid.NewString(),
		Secret:     uuid.NewString(),
	}

	if err := database.SaveGameSession(session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.StartMiniGameResponse{
		SessionID: session.ID,
		Duration:  30,
		Secret:    session.Secret,
	})
}

// FinishMiniGame godoc
// @Summary Завершить мини-игру
// @Tags MiniGame
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body FinishMiniGameRequest true "Game Result"
// @Success 200 {object} map[string]string
// @Router /api/minigame/finish [post]
func FinishMiniGame(c *gin.Context) {
	var req FinishMiniGameRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	session, err := database.GetGameSession(req.SessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "session not found",
		})
		return
	}

	expectedHash := service.BuildGameHash(
		req.SessionID,
		req.Score,
		session.Secret,
	)

	if expectedHash != req.Hash {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "invalid hash",
		})
		return
	}

	duration := time.Now().Unix() - session.StartedAt

	if duration < 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "suspicious result",
		})
		return
	}

	history := models.GameHistory{
		ID:       uuid.NewString(),
		UserName: session.UserName,
		Score:    req.Score,
		PlayedAt: time.Now().Format(time.RFC3339),
	}

	if err := database.DB.Create(&history).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DeleteGameSession(req.SessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var rewardCreated bool
	var bonusGranted bool

	reward := models.Reward{
		ID:       uuid.NewString(),
		Name:     "Супер награда",
		Claimed:  false,
		UserName: session.UserName,
	}

	if err := database.DB.Create(&reward).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rewardCreated = true

	var rewardsCount int64

	if err := database.DB.
		Model(&models.Reward{}).
		Where("user_name = ?", session.UserName).
		Count(&rewardsCount).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rewardsCount > 0 && rewardsCount%4 == 0 {

		bonusHistory := models.GameHistory{
			ID:       uuid.NewString(),
			UserName: session.UserName,
			Score:    500,
			PlayedAt: time.Now().Format(time.RFC3339),
		}

		if err := database.DB.Create(&bonusHistory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		bonusGranted = true
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "game finished",
		"score":          req.Score,
		"reward_created": rewardCreated,
		"bonus_granted":  bonusGranted,
	})
}

// GetMiniGameLeaderboard godoc
// @Summary Лидерборд мини-игры
// @Tags MiniGame
// @Produce json
// @Success 200 {array} MiniGameLeaderboardEntry
// @Router /api/minigame/leaderboard [get]
func GetMiniGameLeaderboard(c *gin.Context) {

	var leaderboard []MiniGameLeaderboardEntry

	err := database.DB.Raw(`
		SELECT
			user_name,
			MAX(score) AS score
		FROM mini_game_sessions
		WHERE finished = true
		GROUP BY user_name
		ORDER BY score DESC
		LIMIT 100
	`).Scan(&leaderboard).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, leaderboard)
}
