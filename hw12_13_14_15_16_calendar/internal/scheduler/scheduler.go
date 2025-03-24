package scheduler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/queue"
)

type Scheduler struct {
	period  time.Duration
	q       *queue.RabbitMq
	storage app.Repository
	ctx     context.Context
	logger  app.Logger
}

func (s *Scheduler) Run() {
	s.logger.Info("scheduler is starting")
	ticker := time.NewTicker(s.period)
	for {
		select {
		case <-ticker.C:
			events := s.getEvents()
			s.Push(events)
		case <-s.ctx.Done():
			s.logger.Info("scheduler is stopping")
			return
		}
	}
}

func (s *Scheduler) getEvents() []storage.Event {
	events, err := s.storage.GetForReminder()

	if err != nil {
		s.logger.Error(err.Error())
		return make([]storage.Event, 0)
	}

	return events
}

func (s *Scheduler) Push(events []storage.Event) {
	for _, v := range events {
		js, err := json.Marshal(v)
		if err != nil {
			s.logger.Error(err.Error())
			continue
		}

		err = s.q.Produce(js, "application/json")
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	s.ctx = ctx

	err := s.storage.Connect(ctx)
	if err != nil {
		return err
	}

	err = s.q.Start(ctx)

	return err
}

func (s *Scheduler) Stop(ctx context.Context) error {
	err := s.storage.Close(ctx)
	if err != nil {
		return err
	}

	err = s.q.Stop(ctx)

	return err
}

func New(period time.Duration, q *queue.RabbitMq, storage app.Repository, logger app.Logger) Scheduler {
	return Scheduler{
		period:  period,
		q:       q,
		storage: storage,
		logger:  logger,
	}
}
