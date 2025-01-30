package repository

import (
	"fmt"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/config"
	memorystorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/repository/memory"
	sqlstorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/repository/sql"
)

func GetStorage(cfg config.Config) (EventRepositoryExtInterface, error) {
	if cfg.StorageType == "memory" {
		rep := memorystorage.New()
		return rep, nil
	}

	if cfg.StorageType == "sql" {
		rep, err := sqlstorage.New(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)
		return rep, err
	}

	return nil, fmt.Errorf("unknown storage type: %s", cfg.StorageType)
}
