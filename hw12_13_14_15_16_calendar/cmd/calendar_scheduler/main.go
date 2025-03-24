package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/logger"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/queue"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/scheduler"
	memorystorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage/memory"
	sqlstorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage/sql"
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

	qu := queue.NewRabbitMq(config.Queue)

	sch := scheduler.New(config.Period, qu, repository, logg)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := sch.Stop(ctx); err != nil {
			logg.Error("failed to stop scheduler: " + err.Error())
		}
	}()

	logg.Info("scheduler is running...")
	err = sch.Start(ctx)
	if err != nil {
		logg.Error(err.Error())
		return
	}

	sch.Run()
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
