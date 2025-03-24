package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/sender"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/queue"

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

	qu := queue.NewRabbitMq(config.Queue)

	send := sender.New(qu, logg)
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctxg, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := send.Stop(ctxg); err != nil {
			logg.Error("failed to stop sender: " + err.Error())
		}
	}()

	err = send.Start(ctx)
	if err != nil {
		logg.Error(err.Error())
		return
	}

	logg.Info("scheduler is running...")
	send.Run()
}
