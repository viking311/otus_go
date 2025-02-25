package app

import (
	"context"
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
