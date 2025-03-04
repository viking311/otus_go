package app

import (
	"context"

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
