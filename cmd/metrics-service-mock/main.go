package main

import (
	"encoding/json"
	"flag"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"math/rand"
	"net"
	"sync"
	"time"
)

var (
	prettyLog = flag.Bool("prettyLog", false, "make logs pretty")
	release   = flag.Bool("release", false, "set logger to info level")
)

func startMockClient() {
	log := logging.GetLogger().WithField("op", "main.startMockClient")

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.WithError(err).Error("Error closing connection")
		}
	}(conn)

	encoder := json.NewEncoder(conn)
	for i := 0; i < 50; i++ {
		metrics := domains.Metrics{
			ID:          uint64(i + 1),
			DeviceID:    100500 + uint64(i),
			Pulse:       rand.Float64() * 100,
			Temperature: 36.0 + rand.Float64()*2,
			LoadCell: domains.LoadCell{
				Output1: rand.Float64() * 50,
				Output2: rand.Float64() * 50,
			},
			MuscleActivity: domains.MuscleActivity{
				Output1: rand.Float64() * 20,
				Output2: rand.Float64() * 20,
			},
		}

		if err := encoder.Encode(metrics); err != nil {
			log.Println("Error encoding JSON:", err)
			return
		}
		log.Println("Sent mock metrics:", metrics)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()

	logging.InitDefaultLogger()
	logging.UpdateByFlags(logging.Flags{PrettyLog: prettyLog, Release: release})

	var wg sync.WaitGroup
	tries := 5
	wg.Add(tries)
	for i := 0; i < tries; i++ {
		go startMockClient()
	}
	wg.Wait()
}
