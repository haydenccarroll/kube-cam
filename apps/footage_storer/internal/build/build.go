package build

import (
	"fmt"
	"os"

	"github.com/haydenccarroll/kube-cam/apps/footage_storer/internal/footagegetter"
)

type Config struct {
	RabbitMQ *footagegetter.RabbitMQ
}

func BuildConfig() (*Config, error) {
	rabbitMQ, err := buildRabbitMQFootageSender()
	if err != nil {
		return nil, fmt.Errorf("failed to build RabbitMQ footage sender: %w", err)
	}

	return &Config{
		RabbitMQ: rabbitMQ,
	}, nil
}

func buildRabbitMQFootageSender() (*footagegetter.RabbitMQ, error) {
	url := os.Getenv("RABBITMQ_URL")
	storePath := os.Getenv("STORE_PATH")
	readQueueName := os.Getenv("READ_QUEUE_NAME")
	return footagegetter.NewRabbitMQ(url, readQueueName, storePath)
}
