package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	t.Run("unknown log level", func(t *testing.T) {
		_, err := New("some strage lavel")

		require.Error(t, err)
	})

	t.Run("correct log level", func(t *testing.T) {
		_, err := New("DEBUG")

		require.NoError(t, err)
	})
}

func TestLogger(t *testing.T) {
	logger, _ := New("INFO")
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
