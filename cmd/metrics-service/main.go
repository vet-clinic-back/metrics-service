package main

import (
	"context"
	"flag"
	"github.com/vet-clinic-back/metrics-service/internal/adapters"
	"github.com/vet-clinic-back/metrics-service/internal/config"
	"github.com/vet-clinic-back/metrics-service/internal/handlers"
	"github.com/vet-clinic-back/metrics-service/internal/services"
	"github.com/vet-clinic-back/metrics-service/internal/storages"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

var (
	prettyLog = flag.Bool("prettyLog", false, "make logs pretty")
	release   = flag.Bool("release", false, "set logger to info level")
)

//  @title      Vet clinic metrics service
//  @version    0.1
//  @description  metrics service

//  @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in              header
// @name            Authorization.
func main() {
	flag.Parse()

	logging.InitDefaultLogger()
	logging.UpdateByFlags(logging.Flags{PrettyLog: prettyLog, Release: release})

	log := logging.GetLogger().WithField("op", "main")
	log.Info("starting metrics-service")

	cfg := config.MustConfigure()

	ad := adapters.NewAdapters(cfg)
	st := storages.MustNew(ad)
	srv := services.MustNew(st)
	h := handlers.New(srv)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	defer stop()

	h.Run(ctx, ad)
}
