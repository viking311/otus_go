package memorystorage

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

type Storage struct {
	events map[string]storage.Event
	mu     sync.RWMutex
}

func (s *Storage) Save(event storage.Event) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if event.ID == "" {
		event.ID = uuid.New().String()
	}

	s.events[event.ID] = event

	return event, nil
}

func (s *Storage) Delete(event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, event.ID)

	return nil
}

func (s *Storage) GetByID(id string) (*storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if ev, ok := s.events[id]; ok {
		return &ev, nil
	}

	return nil, errors.New("event not found")
}

func (s *Storage) GetByUserID(userID int64) (storage.EventList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(storage.EventList, 0)

	for _, ev := range s.events {
		if ev.UserID == userID {
			result = append(result, ev)
		}
	}

	return result, nil
}

func (s *Storage) GetByUserIDAndPeriod(userID int64, dateFrom, dateTo time.Time) (storage.EventList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(storage.EventList, 0)

	for _, ev := range s.events {
		if ev.UserID == userID && ev.DateTime.Unix() >= dateFrom.Unix() && ev.DateTime.Unix() <= dateTo.Unix() {
			result = append(result, ev)
		}
	}

	return result, nil
}

func (s *Storage) GetAll() (storage.EventList, error) {
	eventList := make(storage.EventList, 0, len(s.events))

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
		events: map[string]storage.Event{},
		mu:     sync.RWMutex{},
	}
}
