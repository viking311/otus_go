package main

import (
	"flag"
	"os"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/queue"

	sqlstorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage/sql"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	StorageType string              `yaml:"storageType"`
	Period      time.Duration       `yaml:"period""`
	Logger      logger.Config       `yaml:"logger"`
	DB          sqlstorage.DBConfig `yaml:"db"`
	Queue       queue.Config        `yaml:"queueServer"`
}

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/scheduler_config.yml", "Path to configuration file")
}

func NewConfig() (*Config, error) {
	flag.Parse()

	cfg := Config{}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()
	yamlDecoder := yaml.NewDecoder(file)

	err = yamlDecoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
