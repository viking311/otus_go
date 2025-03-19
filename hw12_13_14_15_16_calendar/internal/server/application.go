package server

import (
	"context"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

type Application interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	GetEvents() storage.EventList
	GetEventByID(id string) *storage.Event
	DeleteEvent(id string)
	SaveEvent(event storage.Event) (*storage.Event, error)
	GetEventsByUserID(userID int64) storage.EventList
	GetEventsByUserIDAndDates(userID int64, dateFrom, dateTo time.Time) storage.EventList
}
