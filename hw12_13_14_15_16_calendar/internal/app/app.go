package app

import (
	"context"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

type App struct {
	storage Repository
	logger  Logger
}

func New(logger Logger, storage Repository) *App {
	return &App{
		logger:  logger,
		storage: storage,
	}
}

func (a *App) Start(ctx context.Context) error {
	return a.storage.Connect(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	return a.storage.Close(ctx)
}

func (a *App) GetEvents() storage.EventList {
	result, err := a.storage.GetAll()

	if err != nil {
		a.logger.Error(err.Error())
		return storage.EventList{}
	}

	return result
}

func (a *App) GetEventById(id string) *storage.Event {
	event, err := a.storage.GetByID(id)
	if err != nil {
		a.logger.Error(err.Error())
		return nil
	}

	return event
}

func (a *App) DeleteEvent(id string) {
	event, err := a.storage.GetByID(id)
	if err != nil {
		a.logger.Error(err.Error())
		return
	}

	err = a.storage.Delete(*event)
	if err != nil {
		a.logger.Error(err.Error())
	}
}

func (a *App) SaveEvent(event storage.Event) (*storage.Event, error) {
	if len(event.Title) == 0 {
		return nil, &FieldValidationError{
			field: "title",
			msg:   "value can't be empty",
		}
	}

	if event.UserID <= 0 {
		return nil, &FieldValidationError{
			field: "userID",
			msg:   "value must be greater than zero",
		}
	}

	if event.DateTime.Unix() <= time.Now().Unix() {
		return nil, &FieldValidationError{
			field: "dateTime",
			msg:   "value must be in the future",
		}
	}

	if event.RemindTime <= 0 {
		return nil, &FieldValidationError{
			field: "RemindTime",
			msg:   "value must be greater than zero",
		}
	}

	newEvent, err := a.storage.Save(event)
	if err != nil {
		return nil, err
	}

	return &newEvent, nil
}
