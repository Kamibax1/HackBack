package service

import (
	"Otbor/internal/database"
	"Otbor/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRewards godoc
// @Summary Инвентарь игрока
// @Tags Rewards
// @Produce json
// @Param username query string true "Username"
// @Success 200 {array} models.Reward
// @Router /api/rewards [get]
func GetRewards(c *gin.Context) {
	username := c.Query("username")

	var rewards []models.Reward

	if err := database.DB.
		Where("user_name = ?", username).
		Find(&rewards).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rewards)
}

type ClaimRewardRequest struct {
	RewardID string `json:"reward_id"`
}

// ClaimReward godoc
// @Summary Получить награду
// @Tags Rewards
// @Accept json
// @Produce json
// @Param request body ClaimRewardRequest true "Reward ID"
// @Success 200 {object} models.Reward
// @Failure 404 {object} map[string]string
// @Router /api/rewards/claim [post]
func ClaimReward(c *gin.Context) {
	var req ClaimRewardRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	var reward models.Reward

	if err := database.DB.
		Where("id = ?", req.RewardID).
		First(&reward).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "reward not found",
		})
		return
	}

	if reward.Claimed {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "reward already claimed",
		})
		return
	}

	reward.Claimed = true

	if err := database.DB.Save(&reward).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reward)
}
