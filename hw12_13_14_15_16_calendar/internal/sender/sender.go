package sender

import (
	"context"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/queue"
)

type Sender struct {
	q      *queue.RabbitMq
	ctx    context.Context
	logger app.Logger
}

func (s *Sender) Run() {
	ch, err := s.q.Consume()
	if err != nil {
		s.logger.Error(err.Error())
		return
	}
	for {
		select {
		case msg := <-ch:
			s.logger.Info("received event" + string(msg.Body))
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *Sender) Start(ctx context.Context) error {
	s.ctx = ctx

	return s.q.Start(ctx)
}

func (s *Sender) Stop(ctx context.Context) error {
	return s.q.Stop(ctx)
}

func New(q *queue.RabbitMq, logger app.Logger) Sender {
	return Sender{
		q:      q,
		logger: logger,
	}
}
