package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vet-clinic-back/metrics-service/internal/delivery/http"
	"github.com/vet-clinic-back/metrics-service/internal/repository/clickhouse"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	// ClickHouse connection
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d",
			viper.GetString("clickhouse.host"),
			viper.GetInt("clickhouse.port"),
		)},
		Auth: clickhouse.Auth{
			Database: viper.GetString("clickhouse.database"),
			Username: viper.GetString("clickhouse.user"),
			Password: viper.GetString("clickhouse.password"),
		},
	})
	if err != nil {
		logger.Fatal(err)
	}

	// Repository
	repo := clickhouse.New(conn, logger)

	// Usecase
	uc := usecase.New(repo, logger)

	// HTTP Server
	router := gin.New()
	router.Use(gin.Recovery())

	// Swagger
	router.GET("/swagger/*any", http.SwaggerHandler())

	// Handlers
	http.NewHttpHandler(router, uc, logger)

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
