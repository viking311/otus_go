package internalhttp

import (
	"encoding/json"
	"net/http"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type GetEventsResponse struct {
	Events storage.EventList `json:"events"`
}

type GetEventsHandler struct {
	app    Application
	logger app.Logger
}

func (ge *GetEventsHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	response := GetEventsResponse{
		Events: ge.app.GetEvents(),
	}

	data, err := json.Marshal(response)
	if err != nil {
		ge.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write(data)
	if err != nil {
		ge.logger.Error(err.Error())
	}
}

func NewGetEventsHandler(app Application, logger app.Logger) *GetEventsHandler {
	return &GetEventsHandler{
		app:    app,
		logger: logger,
	}
}
