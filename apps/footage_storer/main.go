package main

import (
	"context"
	"log"

	"github.com/haydenccarroll/kube-cam/apps/footage_storer/internal/build"
)

func main() {
	config, err := build.BuildConfig()
	if err != nil {
		log.Printf("Failed to build HTTP client: %v", err)
		return
	}
	msgsStored, err := config.RabbitMQ.ReadAndStore(context.Background())
	log.Printf("Successfully stored %d frames\n", msgsStored)
	if err != nil {
		log.Printf("Failed to read and store frames: %v", err)
		return
	}
}
