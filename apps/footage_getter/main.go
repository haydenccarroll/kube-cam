package main

import (
	"context"
	"fmt"
	"log"

	"github.com/haydenccarroll/kube-cam/apps/footage_getter/internal/build"
)

func main() {
	config, err := build.BuildConfig()
	if err != nil {
		log.Printf("Failed to build HTTP client: %v", err)
		return
	}
	continuouslySendFrames(context.Background(), config)
}

func continuouslySendFrames(ctx context.Context, config *build.Config) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			frame, err := config.FrameGetter.Get(ctx)
			if err != nil {
				log.Printf("Failed to get frame: %v", err)
				continue
			}

			err = config.RabbitMQ.PublishToQueue(ctx, "frames", frame)
			if err != nil {
				log.Printf("Failed to publish frame: %v", err)
				continue
			}

			fmt.Println("successfully sent a frame to the queue")
		}
	}
}
