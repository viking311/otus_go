package internal

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/config"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/logger"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/repository"
	internalhttp "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/server/http"
)

type Application struct {
	cfg *config.Config
}

func (app Application) Run() {
	logg, err := logger.New(app.cfg.Logger.Level)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := repository.GetStorage(*app.cfg)
	if err != nil {
		logg.Fatal(err.Error())
	}

	server := internalhttp.NewServer(logg, storage)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

func NewApp() (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	return &Application{
		cfg: cfg,
	}, nil
}
