package handler

import (
	"encoding/json"
	"net/http"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type GetEventHandler struct {
	app    server.Application
	logger app.Logger
}

func (ge *GetEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	event := ge.app.GetEventById(id)

	if event == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(event)
	if err != nil {
		ge.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		ge.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewGetEventById(app server.Application, logger app.Logger) *GetEventHandler {
	return &GetEventHandler{
		app:    app,
		logger: logger,
	}
}
