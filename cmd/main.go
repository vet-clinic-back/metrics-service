package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vet-clinic-back/metrics-service/internal/delivery/http"
	"github.com/vet-clinic-back/metrics-service/internal/repository/postgres"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"
)

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logger.Fatalf("Error reading config: %s", err)
	}

	// PostgreSQL connection
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.host"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.database"),
	)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Repository
	repo := postgres.New(conn, logger)

	// Usecase
	uc := usecase.New(repo, logger)

	// HTTP Server
	router := gin.Default()

	// Handlers
	http.NewHandler(router, uc, logger)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server shutdown error: %s", err)
	}
}
