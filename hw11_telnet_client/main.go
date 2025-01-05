package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

const defaultTimeout = 10

func main() {
	timeout := flag.Duration("timeout", defaultTimeout*time.Second, "timeout connection default 10s")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Port or host not specified")
	}

	host := args[0]
	port := args[1]

	client := NewTelnetClient(net.JoinHostPort(host, port), *timeout, os.Stdin, os.Stdout)

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
		if err := client.Send(); err != nil {
			log.Fatalf("Error reading from channel. Error: %s", err)
		}
	}()

	go func() {
		defer cancel()
		if err := client.Receive(); err != nil {
			log.Fatalf("Error reading from channel. Error: %s", err)
		}
	}()
	<-ctx.Done()
}
