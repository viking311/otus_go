package app

import (
	"context"
)

type App struct {
	storage RepositoryInterface
	logger  Logger
}

func New(logger Logger, storage RepositoryInterface) *App {
	return &App{
		logger:  logger,
		storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

func (a *App) Start(ctx context.Context) error {
	return a.storage.Connect(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	return a.storage.Close(ctx)
}
