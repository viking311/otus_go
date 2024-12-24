package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Success case simple", func(t *testing.T) {
		cmd := []string{"pwd", "-L"}
		returnCode := RunCmd(cmd, Environment{})

		require.Equal(t, 0, returnCode)
	})

	t.Run("Exec err case", func(t *testing.T) {
		cmd := []string{"pwd", "-RRR"}
		returnCode := RunCmd(cmd, Environment{
			"TEST": EnvValue{
				Value: "TEST",
			},
		})

		require.Equal(t, 1, returnCode)
	})
}
