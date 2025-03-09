package storage

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID          string
	Title       string
	UserID      int64
	DateTime    time.Time
	Description string
	Duration    time.Duration
	RemindTime  int64
}

type jsonEvent struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	UserID      int64  `json:"userID"`
	DateTime    int64  `json:"dateTime"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	RemindTime  int64  `json:"remindTime"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonEvent{
		ID:          e.ID,
		Title:       e.Title,
		UserID:      e.UserID,
		DateTime:    e.DateTime.Unix(),
		Description: e.Description,
		Duration:    e.Duration.String(),
		RemindTime:  e.RemindTime,
	})
}

func (e *Event) UnmarshalJSON(data []byte) error {
	stubStruct := jsonEvent{}

	err := json.Unmarshal(data, &stubStruct)
	if err != nil {
		return err
	}

	dr, err := time.ParseDuration(stubStruct.Duration)
	if err != nil {
		return err
	}

	e.ID = stubStruct.ID
	e.Title = stubStruct.Title
	e.UserID = stubStruct.UserID
	e.DateTime = time.Unix(stubStruct.DateTime, 0)
	e.Description = stubStruct.Description
	e.Duration = dr
	e.RemindTime = stubStruct.RemindTime

	return nil
}
