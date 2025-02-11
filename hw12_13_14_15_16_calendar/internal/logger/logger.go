package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	logger *logrus.Logger
}

func New(level string) (*Logger, error) {
	logger := logrus.New()

	intLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(intLevel)

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logger.SetFormatter(formatter)

	return &Logger{
			logger: logger,
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
