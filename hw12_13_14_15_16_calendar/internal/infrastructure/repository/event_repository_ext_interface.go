package repository

import (
	"context"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/domain/repository"
)

type EventRepositoryExtInterface interface {
	repository.EventRepositoryInterface
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}
