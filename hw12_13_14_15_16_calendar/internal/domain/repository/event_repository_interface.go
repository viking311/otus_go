package repository

import (
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/domain/entity"
)

type EventRepositoryInterface interface {
	Save(event entity.Event) (entity.Event, error)
	Delete(event entity.Event) error
	GetByUserID(userID int64) (entity.EventList, error)
	GetByUserIDAndPeriod(userID int64, dateFrom, dateTo time.Time) (entity.EventList, error)
	GetAll() (entity.EventList, error)
}
