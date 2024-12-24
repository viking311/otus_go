package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	testDir, err := os.MkdirTemp("", "empty_dir")
	if err != nil {
		t.Fatal("can't create temp dir: ", err)
	}

	tests := []struct {
		name     string
		envDir   string
		expected Environment
		isError  bool
	}{
		{
			name:   "success test",
			envDir: "testdata/env",
			expected: Environment{
				"BAR":   EnvValue{Value: "bar", NeedRemove: false},
				"EMPTY": EnvValue{Value: "", NeedRemove: true},
				"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
				"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
				"UNSET": EnvValue{Value: "", NeedRemove: true},
			},
			isError: false,
		},
		{
			name:     "success test with empty",
			envDir:   testDir,
			expected: Environment{},
			isError:  false,
		},
		{
			name:     "non exists directory",
			envDir:   "testdata/testdata/",
			expected: Environment{},
			isError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env, err := ReadDir(tt.envDir)
			if tt.isError {
				require.Error(t, err)
			} else {
				require.Equal(t, tt.expected, env)
				require.NoError(t, err)
			}
		})
	}
}
