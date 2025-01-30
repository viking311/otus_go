package config

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	StorageType string     `toml:"storage_type"`
	Logger      LoggerConf `toml:"logger"`
	DB          DBConfig   `toml:"db"`
}

type LoggerConf struct {
	Level string `toml:"level"`
}

type DBConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"db_name"`
}

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
}

func NewConfig() (*Config, error) {
	flag.Parse()

	var cfg Config
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		errC := file.Close()
		if errC != nil {
			log.Println(errC)
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
