package footagesender

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
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
	}, nil
}

func (r *RabbitMQ) Close() {
	r.Channel.Close()
	r.Connection.Close()
}

func (r *RabbitMQ) PublishToQueue(ctx context.Context, queueName string, data []byte) error {
	q, err := r.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = r.Channel.PublishWithContext(
		ctx,    // context
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "image/png",
			Body:        data,
		})
	if err != nil {
		return err
	}

	return nil
}
