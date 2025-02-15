package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vet-clinic-back/metrics-service/docs"
	"github.com/vet-clinic-back/metrics-service/internal/config"
	"github.com/vet-clinic-back/metrics-service/internal/handler"
	"github.com/vet-clinic-back/metrics-service/internal/repository/postgres"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"
)

// @title Health Monitoring API
// @version 1.0
// @description API для сбора и анализа медицинских показателей
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@healthmonitor.com

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

	// Database setup
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
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

	// Dependency setup
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

	// Добавляем маршрут для Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpHandler := handler.NewHTTPHandler(uc, logger)
	router.GET("/metrics", httpHandler.GetMetrics)
	router.POST("/metrics", httpHandler.SaveMetrics)

	if err := router.Run(cfg.HTTP.Addr); err != nil {
		logger.Fatal("HTTP server failed:", err)
	}
}
