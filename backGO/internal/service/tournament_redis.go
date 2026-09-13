package service

import (
	"Otbor/internal/database"
	"Otbor/internal/models"
	"encoding/json"
	"time"
)

type ActiveTournament struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartedAt time.Time `json:"started_at"`
	EndsAt    time.Time `json:"ends_at"`
}

func GetActiveTournament() (*ActiveTournament, error) {
	data, err := database.RDB.Get(
		database.Ctx,
		"tournament:active",
	).Result()

	if err != nil {
		return nil, err
	}

	var tournament ActiveTournament

	if err := json.Unmarshal(
		[]byte(data),
		&tournament,
	); err != nil {
		return nil, err
	}

	return &tournament, nil
}

func SetActiveTournament(
	tournament ActiveTournament,
) error {

	data, err := json.Marshal(
		tournament,
	)

	if err != nil {
		return err
	}

	return database.RDB.Set(
		database.Ctx,
		"tournament:active",
		data,
		0,
	).Err()
}

func TournamentWatcher() {
	for {
		active, err := GetActiveTournament()
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		if time.Now().Before(active.EndsAt) {
			println("not ended yet")
			time.Sleep(10 * time.Second)
			continue
		}

		var winner *string

		top, err := database.GetTournamentTop(active.ID, 1)
		if err == nil && len(top) > 0 {
			w := top[0].Member.(string)
			winner = &w
		}

		tournament := models.Tournament{
			Name:      active.Name,
			StartedAt: active.StartedAt,
			EndedAt:   active.EndsAt,
			Winner:    winner,
		}

		if err := database.DB.Create(&tournament).Error; err != nil {
			println("db error:", err.Error())
			time.Sleep(10 * time.Second)
			continue
		}

		database.RDB.Del(database.Ctx, "tournament:active")
		database.RDB.Del(database.Ctx, "tournament:"+active.ID+":leaderboard")

		time.Sleep(10 * time.Second)
	}
}
