package storage

import "time"

type Event struct {
	ID          string
	Title       string
	UserID      int64
	DateTime    time.Time
	Description string
	Duration    time.Duration
	RemindTime  int64
}
