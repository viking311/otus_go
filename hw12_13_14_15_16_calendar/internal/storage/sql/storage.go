package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // подключаем драйвер
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

type Storage struct {
	db  *sql.DB
	ctx context.Context
}

func New(host string, port int, user, password, dbName string) (*Storage, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Connect(ctx context.Context) error {
	err := s.db.PingContext(ctx)
	if err != nil {
		return err
	}

	s.ctx = ctx

	return nil
}

func (s *Storage) Close(_ context.Context) error {
	err := s.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Save(event storage.Event) (storage.Event, error) {
	if event.ID == "" {
		event, err := s.insert(event)
		return event, err
	}

	err := s.update(event)

	return event, err
}

func (s *Storage) Delete(event storage.Event) error {
	_, err := s.db.ExecContext(s.ctx, "DELETE FROM events where id=$1", event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetByUserID(userID int64) (storage.EventList, error) {
	eventList := make(storage.EventList, 0)

	rows, err := s.db.QueryContext(
		s.ctx,
		"SELECT id,title,description,event_time,duration,remind_time,user_id FROM events WHERE user_id=$1",
		userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var desc, duration sql.NullString
		var rTime sql.NullInt64
		var title, id string
		var eventTime time.Time
		var userID int64

		err = rows.Scan(
			&id,
			&title,
			&desc,
			&eventTime,
			&duration,
			&rTime,
			&userID,
		)
		if err != nil {
			return nil, err
		}

		event := storage.Event{
			ID:       id,
			Title:    title,
			DateTime: eventTime,
			UserID:   userID,
		}

		if desc.Valid {
			event.Description = desc.String
		}

		if duration.Valid {
			if d, err := time.ParseDuration(duration.String); err == nil {
				event.Duration = d
			}
		}

		if rTime.Valid {
			event.RemindTime = rTime.Int64
		}

		eventList = append(eventList, event)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return eventList, nil
}

func (s *Storage) GetByUserIDAndPeriod(userID int64, dateFrom, dateTo time.Time) (storage.EventList, error) {
	eventList := make(storage.EventList, 0)

	rows, err := s.db.QueryContext(
		s.ctx,
		`SELECT 
    				id,
    				title,
    				description,
    				event_time,
    				duration,
    				remind_time,
    				user_id 
				FROM events 
				WHERE user_id=$1 and event_date between $2 AND $3`,
		userID,
		dateFrom.Format(time.RFC3339),
		dateTo.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var desc, duration sql.NullString
		var rTime sql.NullInt64
		var title, id string
		var eventTime time.Time
		var userID int64

		err = rows.Scan(
			&id,
			&title,
			&desc,
			&eventTime,
			&duration,
			&rTime,
			&userID,
		)
		if err != nil {
			return nil, err
		}

		event := storage.Event{
			ID:       id,
			Title:    title,
			DateTime: eventTime,
			UserID:   userID,
		}

		if desc.Valid {
			event.Description = desc.String
		}

		if duration.Valid {
			if d, err := time.ParseDuration(duration.String); err == nil {
				event.Duration = d
			}
		}

		if rTime.Valid {
			event.RemindTime = rTime.Int64
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return eventList, nil
}

func (s *Storage) GetAll() (storage.EventList, error) {
	eventList := make(storage.EventList, 0)

	rows, err := s.db.QueryContext(
		s.ctx,
		`SELECT 
    				id,
    				title,
    				description,
    				event_time,
    				duration,
    				remind_time,
    				user_id 
				FROM events`)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var desc, duration sql.NullString
		var rTime sql.NullInt64
		var title, id string
		var eventTime time.Time
		var userID int64

		err = rows.Scan(
			&id,
			&title,
			&desc,
			&eventTime,
			&duration,
			&rTime,
			&userID,
		)
		if err != nil {
			return nil, err
		}

		event := storage.Event{
			ID:       id,
			Title:    title,
			DateTime: eventTime,
			UserID:   userID,
		}

		if desc.Valid {
			event.Description = desc.String
		}

		if duration.Valid {
			if d, err := time.ParseDuration(duration.String); err == nil {
				event.Duration = d
			}
		}

		if rTime.Valid {
			event.RemindTime = rTime.Int64
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return eventList, nil
}

func (s *Storage) insert(event storage.Event) (storage.Event, error) {
	err := s.db.QueryRowContext(
		s.ctx,
		`INSERT INTO events(title, description, event_time, duration, remind_time, user_id) 
				VALUES($1,$2,$3,$4,$5,$6) RETURNING id`,
		event.Title,
		event.Description,
		event.DateTime.Format(time.RFC3339),
		event.Duration.String(),
		event.RemindTime,
		event.UserID,
	).Scan(&event.ID)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (s *Storage) update(event storage.Event) error {
	_, err := s.db.ExecContext(
		s.ctx,
		`UPDATE events SET 
                  title=$1, 
                  description=$2, 
                  datetime=$3, 
                  duration=$4, 
                  remind_time=$5, 
                  user_id=$6 
              WHERE id=$7`,
		event.Title,
		event.Description,
		event.DateTime.Format(time.RFC3339),
		event.Duration.String(),
		event.RemindTime,
		event.UserID,
		event.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
