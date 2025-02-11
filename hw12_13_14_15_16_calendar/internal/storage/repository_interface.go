package storage

import (
	"context"
	"time"
)

type RepositoryInterface interface {
	Save(event Event) (Event, error)
	Delete(event Event) error
	GetByUserID(userID int64) (EventList, error)
	GetByUserIDAndPeriod(userID int64, dateFrom, dateTo time.Time) (EventList, error)
	GetAll() (EventList, error)
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}
