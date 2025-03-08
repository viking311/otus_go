package app

import (
	"context"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

type Repository interface {
	Save(event storage.Event) (storage.Event, error)
	Delete(event storage.Event) error
	GetByID(id string) (*storage.Event, error)
	GetByUserID(userID int64) (storage.EventList, error)
	GetByUserIDAndPeriod(userID int64, dateFrom, dateTo time.Time) (storage.EventList, error)
	GetAll() (storage.EventList, error)
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}
