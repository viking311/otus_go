package main

import (
	"flag"
	"log"

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

	logg.Info("scheduler is running...")
}
