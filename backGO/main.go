package main

import (
	"crypto/rand"
	"encoding/base64"

	_ "Otbor/docs"
	"Otbor/internal/database"
	"Otbor/internal/handlers"
	"Otbor/internal/middleware"
	"Otbor/internal/models"
	"Otbor/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SeedAdmin() {
	var admin models.User

	err := database.DB.Where("user_role = ?", "admin").
		First(&admin).Error

	if err == nil {
		return
	}

	salt := make([]byte, 16)
	_, _ = rand.Read(salt)

	admin = models.User{
		UserName:     "admin",
		UserEmail:    "admin@localhost",
		UserRole:     "admin",
		PasswordHash: service.HashPassword("admin123", salt),
		Salt:         base64.StdEncoding.EncodeToString(salt),
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		panic(err)
	}
}

// @title User API
// @version 1.0
// @description Simple user service
// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	database.ConnectDB()
	database.ConnectRedis()
	go service.TournamentWatcher()
	SeedAdmin()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
	}))

	router.POST("/api/user/login", service.Login)
	router.POST("/api/user/register", service.Register)
	router.POST("/api/user/update", service.Update)
	router.DELETE("/api/user/delete", service.Delete)
	router.GET("/api/user/profile/:username", service.Profile)
	router.GET("/api/tournament", handlers.GetTournaments)
	router.GET("/api/latest_tournaments", handlers.GetLatestTournaments)
	router.GET("/api/tournament/top", service.GetLeaderBoard)
	router.GET("/api/rewards", service.GetRewards)
	router.POST("/api/rewards/claim", service.ClaimReward)

	admin := router.Group("/api/admin")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
	)
	{
		admin.GET("/users", handlers.AdminGetUsers)
		admin.GET("/users/:username", handlers.AdminGetUser)
		admin.PUT("/users/role", handlers.AdminChangeRole)
		admin.DELETE("/users/:username", handlers.AdminDeleteUser)
		admin.GET("/tournaments", handlers.GetTournaments)
		admin.POST("/tournaments", handlers.AdminCreateTournament)
		admin.DELETE("/tournaments/:id", handlers.AdminDeleteTournament)

		admin.POST("/rewards", handlers.AdminCreateReward)
		admin.GET("/rewards", handlers.AdminGetRewards)

		admin.GET("/games", handlers.AdminGetGames)
	}
	miniGame := router.Group("/api/minigame")
	miniGame.Use(middleware.AuthMiddleware())
	{
		miniGame.POST("/start", handlers.StartMiniGame)
		miniGame.POST("/finish", handlers.FinishMiniGame)
	}

	router.GET(
		"/api/minigame/leaderboard",
		handlers.GetMiniGameLeaderboard,
	)

	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8000")
}
