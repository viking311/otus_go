package memorystorage

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

func TestStorage_ConnectClose(t *testing.T) {
	t.Run("connect", func(t *testing.T) {
		memoryStorage := New()
		ctx := context.Background()

		err := memoryStorage.Connect(ctx)

		require.NoError(t, err)
	})

	t.Run("close", func(t *testing.T) {
		memoryStorage := New()
		ctx := context.Background()

		err := memoryStorage.Close(ctx)

		require.NoError(t, err)
	})
}

func TestStorage_Save(t *testing.T) {
	t.Run("add new event", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		eventActual, err := memoryStorage.Save(event)

		require.NoError(t, err)
		require.Equal(t, len(memoryStorage.events), 1)
		require.NotEmpty(t, eventActual.ID)

		eventActual.ID = ""
		require.Equal(t, event, eventActual)
	})

	t.Run("update event", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		event, _ = memoryStorage.Save(event)

		event.UserID = 2

		eventActual, err := memoryStorage.Save(event)

		require.NoError(t, err)
		require.Equal(t, len(memoryStorage.events), 1)
		require.Equal(t, event, eventActual)
	})
}

func TestStorage_Delete(t *testing.T) {
	t.Run("delete event", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		event, _ = memoryStorage.Save(event)

		err := memoryStorage.Delete(event)

		require.NoError(t, err)
		require.Equal(t, 0, len(memoryStorage.events))
	})
}

func TestStorage_GetById(t *testing.T) {
	t.Run("get by id with empty result", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		_, _ = memoryStorage.Save(event)

		_, ok := memoryStorage.GetByID("unknown ID")

		require.Equal(t, false, ok)
	})

	t.Run("get by id success", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		event, _ = memoryStorage.Save(event)

		res, ok := memoryStorage.GetByID(event.ID)

		require.Equal(t, true, ok)
		require.Equal(t, event, res)
	})
}

func TestStorage_GetAll(t *testing.T) {
	t.Run("get all events", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		event, _ = memoryStorage.Save(event)

		events, err := memoryStorage.GetAll()

		require.NoError(t, err)
		require.Equal(t, len(memoryStorage.events), len(events))
		require.Equal(t, event, events[0])
	})
}

func TestStorage_GetByID(t *testing.T) {
	t.Run("get by user id with empty result", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		_, _ = memoryStorage.Save(event)

		events, err := memoryStorage.GetByUserID(2)

		require.NoError(t, err)
		require.Equal(t, 0, len(events))
	})

	t.Run("get by user id success", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		event, _ = memoryStorage.Save(event)

		events, err := memoryStorage.GetByUserID(event.UserID)

		require.NoError(t, err)
		require.Equal(t, 1, len(events))
		require.Equal(t, event, events[0])
	})
}

func TestStorage_GetByUserIDAndPeriod(t *testing.T) {
	t.Run("get by user id and dates (unknow user)", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		_, _ = memoryStorage.Save(event)
		dateFrom := time.Now().Truncate(1 * time.Hour)
		dateTo := time.Now().Add(2 * time.Hour)
		events, err := memoryStorage.GetByUserIDAndPeriod(2, dateFrom, dateTo)

		require.NoError(t, err)
		require.Equal(t, 0, len(events))
	})

	t.Run("get by user id and dates (no events in period)", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		_, _ = memoryStorage.Save(event)
		dateFrom := time.Now().Add(1 * time.Hour)
		dateTo := time.Now().Add(2 * time.Hour)
		events, err := memoryStorage.GetByUserIDAndPeriod(1, dateFrom, dateTo)

		require.NoError(t, err)
		require.Equal(t, 0, len(events))
	})

	t.Run("get by user id and dates success", func(t *testing.T) {
		memoryStorage := New()

		event := storage.Event{
			Title:       "title",
			UserID:      1,
			DateTime:    time.Now(),
			Description: "description",
			Duration:    1 * time.Hour,
			RemindTime:  15,
		}

		event, _ = memoryStorage.Save(event)
		dateFrom := time.Now().Truncate(1 * time.Hour)
		dateTo := time.Now().Add(2 * time.Hour)
		events, err := memoryStorage.GetByUserIDAndPeriod(1, dateFrom, dateTo)

		require.NoError(t, err)
		require.Equal(t, 1, len(events))
		require.Equal(t, event, events[0])
	})
}
