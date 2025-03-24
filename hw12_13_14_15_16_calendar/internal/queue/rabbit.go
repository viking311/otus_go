package queue

import (
	"context"
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ErrNotConnected = errors.New("connection isn't exists")

type RabbitMq struct {
	ctx  context.Context
	cfg  Config
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func (r *RabbitMq) Start(ctx context.Context) error {
	r.ctx = ctx
	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%s/", r.cfg.User, r.cfg.Password, r.cfg.Host, r.cfg.Port),
	)
	if err != nil {
		return err
	}
	r.conn = conn

	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	r.ch = ch

	q, err := ch.QueueDeclare(
		r.cfg.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	r.q = q

	return nil
}

func (r *RabbitMq) Stop(_ context.Context) error {
	err := r.conn.Close()
	if err != nil {
		return err
	}

	err = r.ch.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMq) Produce(msg []byte, contentType string) error {
	if r.ch == nil {
		return ErrNotConnected
	}

	if contentType == "" {
		contentType = "application/json"
	}

	return r.ch.PublishWithContext(
		r.ctx,
		"",
		r.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        msg,
		},
	)
}

func (r *RabbitMq) Consume() (<-chan amqp.Delivery, error) {
	if r.ch == nil {
		return nil, ErrNotConnected
	}

	messages, consumeErr := r.ch.Consume(
		r.q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if consumeErr != nil {
		return nil, consumeErr
	}

	return messages, nil
}

func NewRabbitMq(cfg Config) *RabbitMq {
	return &RabbitMq{
		cfg: cfg,
	}
}
