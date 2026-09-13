package database

import (
	"Otbor/internal/models"
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string

	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func LoadConfig() Config {
	_ = godotenv.Load()

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),

		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDB,
	}
}

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:         "157.228.191.29:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	if err := RDB.Ping(Ctx).Err(); err != nil {
		panic("redis connection failed: " + err.Error())
	}

	println("redis connected")
}

func GetTournamentTop(
	tournamentID string,
	limit int64,
) ([]redis.Z, error) {

	key := "tournament:" + tournamentID + ":leaderboard"

	return RDB.ZRevRangeWithScores(
		Ctx,
		key,
		0,
		limit-1,
	).Result()
}

func AddTournamentScore(
	tournamentID string,
	username string,
	score int,
) error {

	key := "tournament:" + tournamentID + ":leaderboard"

	return RDB.ZIncrBy(
		Ctx,
		key,
		float64(score),
		username,
	).Err()
}

func SaveGameSession(session models.ActiveGameSession) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return RDB.Set(
		Ctx,
		"session:"+session.ID,
		data,
		30*time.Minute,
	).Err()
}

func GetGameSession(id string) (*models.ActiveGameSession, error) {
	val, err := RDB.Get(
		Ctx,
		"session:"+id,
	).Result()

	if err != nil {
		return nil, err
	}

	var session models.ActiveGameSession

	if err := json.Unmarshal([]byte(val), &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func DeleteGameSession(id string) error {
	return RDB.Del(
		Ctx,
		"session:"+id,
	).Err()
}
