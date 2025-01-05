package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	config := GetConfig()

	client := NewTelnetClient(config.Address, config.Timeout, os.Stdin, os.Stdout)

	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		errC := client.Close()
		if errC != nil {
			log.Println(errC)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		defer cancel()

		errS := client.Send()
		if errS != nil {
			log.Fatal(errS)
		}
	}()

	go func() {
		defer cancel()

		errR := client.Receive()
		if errR != nil {
			log.Fatal(errR)
		}
	}()

	<-ctx.Done()
}
