package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	memorystorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage/memory"
	sqlstorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage/sql"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/logger"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.New(config.Logger)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		errClose := logg.Close()
		if errClose != nil {
			log.Fatal(errClose)
		}
	}()

	repository, err := getStorage(config.StorageType, config.DB)
	if err != nil {
		logg.Error(err.Error())
		return
	}

	fmt.Println(repository)

	logg.Info("scheduler is running...")
}

func getStorage(storageType string, cfg sqlstorage.DBConfig) (app.Repository, error) {
	if storageType == "memory" {
		rep := memorystorage.New()
		return rep, nil
	}

	if storageType == "sql" {
		rep, err := sqlstorage.New(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
		return rep, err
	}

	return nil, fmt.Errorf("unknown storage type: %s", storageType)
}
