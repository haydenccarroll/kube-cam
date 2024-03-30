package footagegetter

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
	QueueName  string
	DirPath    string
}

func NewRabbitMQ(url, queueName, dirPath string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
		QueueName:  queueName,
		DirPath:    dirPath,
	}, nil
}

func (r *RabbitMQ) Close() {
	_ = r.Channel.Close()
	_ = r.Connection.Close()
}

func (r *RabbitMQ) ReadAndStore(ctx context.Context) (int, error) {
	msgs, err := r.Channel.Consume(
		r.QueueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return 0, err
	}

	msgsStored := 0
	for {
		select {
		case <-ctx.Done():
			return msgsStored, nil
		case msg, ok := <-msgs:
			if !ok {
				return msgsStored, nil
			}
			timeNow := time.Now()
			timeStr := timeNow.Format("2006-01-02") + "-" + fmt.Sprintf("%d", timeNow.Unix())
			fileName := fmt.Sprintf("%s/%s.png", r.DirPath, timeStr)
			err := os.WriteFile(fileName, msg.Body, 0644)
			if err != nil {
				log.Print(err)
			} else {
				msgsStored++
			}
		}
	}
}
