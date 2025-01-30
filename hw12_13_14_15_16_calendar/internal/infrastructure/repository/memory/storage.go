package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/domain/entity"
)

type Storage struct {
	events map[string]entity.Event
	mu     sync.RWMutex
}

func (s *Storage) Save(event entity.Event) (entity.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if event.ID == "" {
		event.ID = uuid.New().String()
	}

	s.events[event.ID] = event

	return event, nil
}

func (s *Storage) Delete(event entity.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, event.ID)

	return nil
}

func (s *Storage) GetByID(id string) (entity.Event, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.events[id]; ok {
		return s.events[id], true
	}

	return entity.Event{}, false
}

func (s *Storage) GetByUserID(userID int64) (entity.EventList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(entity.EventList, 0)

	for _, ev := range s.events {
		if ev.UserID == userID {
			result = append(result, ev)
		}
	}

	return result, nil
}

func (s *Storage) GetByUserIDAndPeriod(userID int64, dateFrom, dateTo time.Time) (entity.EventList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(entity.EventList, 0)

	for _, ev := range s.events {
		if ev.UserID == userID && ev.DateTime.Unix() >= dateFrom.Unix() && ev.DateTime.Unix() <= dateTo.Unix() {
			result = append(result, ev)
		}
	}

	return result, nil
}

func (s *Storage) GetAll() (entity.EventList, error) {
	eventList := make(entity.EventList, 0, len(s.events))

	for _, v := range s.events {
		eventList = append(eventList, v)
	}

	return eventList, nil
}

func (s *Storage) Connect(_ context.Context) error {
	return nil
}

func (s *Storage) Close(_ context.Context) error {
	return nil
}

func New() *Storage {
	return &Storage{
		events: map[string]entity.Event{},
		mu:     sync.RWMutex{},
	}
}
