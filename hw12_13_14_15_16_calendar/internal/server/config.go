package server

import "time"

type HTTPServerConfig struct {
	BindAddress string        `yaml:"bindAddress"`
	BindPort    string        `yaml:"bindPort"`
	Timeout     time.Duration `yaml:"timeout"`
}

type GRPCServerConfig struct {
	BindAddress string `yaml:"bindAddress"`
	BindPort    string `yaml:"bindPort"`
}
