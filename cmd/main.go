package main

import (
	"database/sql"
	"fmt"

	"github.com/vet-clinic-back/metrics-service/internal/config"
	"github.com/vet-clinic-back/metrics-service/internal/handler"
	"github.com/vet-clinic-back/metrics-service/internal/repository/postgres"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/vet-clinic-back/metrics-service/docs"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Health Metrics API
// @version 1.0
// @description API for health monitoring system
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@healthmetrics.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Config error:", err)
	}

	// Database connection
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode,
	)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		logger.Fatal("DB connection error:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		logger.Fatal("DB ping error:", err)
	}

	dbx := sqlx.NewDb(db, "pgx")

	// Repository and UseCase
	repo := postgres.NewSensorDataRepo(dbx)
	uc := usecase.NewSensorDataUseCase(repo)

	// TCP Server
	tcpHandler := handler.NewTCPHandler(uc, logger)
	go func() {
		if err := tcpHandler.Start(cfg.TCP.Addr); err != nil {
			logger.Fatal("TCP server failed:", err)
		}
	}()

	// HTTP Server
	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpHandler := handler.NewHTTPHandler(uc, logger)
	router.GET("/metrics", httpHandler.GetMetrics)

	if err := router.Run(cfg.HTTP.Addr); err != nil {
		logger.Fatal("HTTP server failed:", err)
	}
}
