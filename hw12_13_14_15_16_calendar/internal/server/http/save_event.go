package internalhttp

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type SaveEventResponse struct {
	Status string         `json:"status"`
	Msg    string         `json:"msg"`
	Event  *storage.Event `json:"event"`
}

type SaveEventHandler struct {
	app    Application
	logger app.Logger
}

func (se *SaveEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	event := storage.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		se.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	savedEvent, err := se.app.SaveEvent(event)
	if err != nil {
		var verr *app.FieldValidationError
		if errors.As(err, &verr) {
			response := SaveEventResponse{
				Status: "error",
				Msg:    verr.Error(),
			}
			err := json.NewEncoder(w).Encode(&response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		se.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := SaveEventResponse{
		Status: "ok",
		Event:  savedEvent,
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		se.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func NewSaveEventHandler(app Application, logger app.Logger) *SaveEventHandler {
	return &SaveEventHandler{
		app:    app,
		logger: logger,
	}
}
