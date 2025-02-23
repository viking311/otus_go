package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	t.Run("unknown log level", func(t *testing.T) {
		_, err := New(Config{Level: "some strage lavel", FileName: "/tmp/calendar.log"})

		require.Error(t, err)
	})

	t.Run("correct log level", func(t *testing.T) {
		_, err := New(Config{Level: "DEBUG", FileName: "/tmp/calendar.log"})

		require.NoError(t, err)
	})
}

func TestLogger(t *testing.T) {
	logger, _ := New(Config{Level: "INFO", FileName: "/tmp/calendar.log"})
	msg := "test message"

	t.Run("write info", func(t *testing.T) {
		out := &bytes.Buffer{}
		logger.logger.Out = out

		logger.Info(msg)
		require.Contains(t, out.String(), msg)
	})

	t.Run("skip message", func(t *testing.T) {
		out := &bytes.Buffer{}
		logger.logger.Out = out

		logger.Debug(msg)
		require.Empty(t, out.String())
	})
}
