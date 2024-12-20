package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name   string
		from   string
		to     string
		offset int64
		limit  int64
	}{
		{
			name:   "Empty from",
			from:   "",
			to:     "out.txt",
			offset: 0,
			limit:  0,
		},
		{
			name:   "Empty to",
			from:   "testdata/input.txt",
			to:     "",
			offset: 0,
			limit:  0,
		},
		{
			name:   "Negative offset",
			from:   "testdata/input.txt",
			to:     "out.txt",
			offset: -10,
			limit:  0,
		},
		{
			name:   "Negative limit",
			from:   "testdata/input.txt",
			to:     "out.txt",
			offset: 0,
			limit:  -1,
		},
		{
			name:   "From file isn't exists",
			from:   "filename",
			to:     "out.txt",
			offset: 0,
			limit:  0,
		},
		{
			name:   "Unsuported file",
			from:   "/dev/urandom",
			to:     "out.txt",
			offset: 0,
			limit:  0,
		},
		{
			name:   "ErrUnsupportedFile",
			from:   "testdata/input.txt",
			to:     "out.txt",
			offset: 1000000,
			limit:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.from, tt.to, tt.offset, tt.limit)
			require.Error(t, err)
		})
	}
}
