package service

import (
	"Otbor/internal/database"
	"Otbor/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGlobalHistory godoc
// @Summary История всех игр
// @Tags Games
// @Produce json
// @Success 200 {array} models.GameHistory
// @Router /api/games/history/global [get]
func GetGlobalHistory(c *gin.Context) {
	var history []models.GameHistory

	if err := database.DB.
		Order("played_at desc").
		Find(&history).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, history)
}

// GetTournament godoc
// @Summary Текущий турнир
// @Tags Tournament
// @Produce json
// @Success 200 {object} models.Tournament
// @Router /api/tournament [get]
func GetTournament(c *gin.Context) {
	tournament, err := GetActiveTournament()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no active tournament",
		})
		return
	}

	c.JSON(http.StatusOK, tournament)
}

// GetTournamentTop godoc
// @Summary Топ игроков турнира
// @Tags Tournament
// @Produce json
// @Success 200 {array} models.LeaderboardEntry
// @Router /api/tournament/top [get]
func GetLeaderBoard(c *gin.Context) {
	tournament, err := GetActiveTournament()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no active tournament",
		})
		return
	}

	top, err := database.GetTournamentTop(
		tournament.ID,
		100,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var leaderboard []models.LeaderboardEntry

	for _, player := range top {
		leaderboard = append(
			leaderboard,
			models.LeaderboardEntry{
				UserName: player.Member.(string),
				Score:    int(player.Score),
			},
		)
	}

	c.JSON(http.StatusOK, leaderboard)
}
