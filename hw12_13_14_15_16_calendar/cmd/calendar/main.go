package main

import (
	"flag"
	"log"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	app, err := internal.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
