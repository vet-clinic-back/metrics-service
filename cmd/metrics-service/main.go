package main

import (
	"flag"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
)

var (
	prettyLog = flag.Bool("prettyLog", false, "make logs pretty")
	release   = flag.Bool("release", false, "set logger to info level")
)

func main() {
	flag.Parse()

	logging.InitDefaultLogger()
	logging.UpdateByFlags(logging.Flags{PrettyLog: prettyLog, Release: release})

	log := logging.GetLogger().WithField("op", "main")
	log.Info("starting metrics-service")

}
