package app

import (
	"context"

	"github.com/viking311/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	storage.RepositoryInterface
}

type Logger interface { // TODO
}

type Storage interface { // TODO
}

func New(logger Logger, storage storage.RepositoryInterface) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
