package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Level    string `yaml:"level"`
	FileName string `yaml:"fileName"`
}

type Logger struct {
	logger *logrus.Logger
	out    io.Closer
}

func New(cfg Config) (*Logger, error) {
	logger := logrus.New()

	intLevel, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(intLevel)

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logger.SetFormatter(formatter)

	file, err := os.OpenFile(cfg.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}
	logger.Out = file

	return &Logger{
			logger: logger,
			out:    file,
		},
		nil
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l Logger) Print(msg string) {
	l.logger.Print(msg)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l Logger) Panic(msg string) {
	l.logger.Panic(msg)
}

func (l Logger) Close() error {
	return l.out.Close()
}
