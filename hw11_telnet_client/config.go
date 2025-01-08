package main

import (
	"flag"
	"log"
	"net"
	"time"
)

type Config struct {
	Timeout time.Duration
	Address string
}

const defaultTimeout = 10 * time.Second

func GetConfig() Config {
	var cfg Config
	var host, port string
	flag.StringVar(&host, "host", "localhost", "host to connect to")
	flag.StringVar(&port, "port", "4242", "port to connect to")
	flag.DurationVar(&cfg.Timeout, "timeout", defaultTimeout, "timeout")
	flag.Parse()

	if host == "" {
		log.Fatal("host is empty")
	}

	if port == "" {
		log.Fatal("port is empty")
	}

	cfg.Address = net.JoinHostPort(host, port)

	return cfg
}
