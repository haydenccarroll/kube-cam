package build

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/haydenccarroll/kube-cam/apps/footage_getter/internal/framegetter"
	"github.com/haydenccarroll/kube-cam/apps/footage_getter/internal/framesender"
)

type Config struct {
	RabbitMQ    *framesender.RabbitMQ
	FrameGetter *framegetter.FrameGetterHTTPCatImpl
}

func BuildConfig() (*Config, error) {
	httpClient, err := buildHTTPClient()
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	rabbitMQ, err := buildRabbitMQFootageSender()
	if err != nil {
		return nil, fmt.Errorf("failed to build RabbitMQ footage sender: %w", err)
	}

	return &Config{
		RabbitMQ:    rabbitMQ,
		FrameGetter: buildFrameGetter(httpClient),
	}, nil
}

func buildHTTPClient() (*http.Client, error) {
	timeout := os.Getenv("HTTP_CLIENT_TIMEOUT")
	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP_CLIENT_TIMEOUT: %w", err)
	}

	return &http.Client{
		Timeout: timeoutDuration,
	}, nil
}

func buildRabbitMQFootageSender() (*framesender.RabbitMQ, error) {
	url := os.Getenv("RABBITMQ_URL")
	queueName := os.Getenv("WRITE_QUEUE_NAME")
	return framesender.NewRabbitMQ(url, queueName)
}

func buildFrameGetter(httpClient *http.Client) *framegetter.FrameGetterHTTPCatImpl {
	url := os.Getenv("FRAME_GETTER_URL")
	return framegetter.NewFrameGetterHTTPCatImpl(httpClient, url)
}
