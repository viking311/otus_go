package main

import (
	"flag"
	"os"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/logger"
	sqlstorage "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage/sql"
	"gopkg.in/yaml.v3"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	StorageType string                  `yaml:"storageType"`
	Logger      logger.Config           `yaml:"logger"`
	DB          sqlstorage.DBConfig     `yaml:"db"`
	HTTPServer  server.HTTPServerConfig `yaml:"httpServer"`
}

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yml", "Path to configuration file")
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
