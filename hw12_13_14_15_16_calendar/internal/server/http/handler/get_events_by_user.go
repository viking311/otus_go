package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

type getEventsByUserResponse struct {
	Events storage.EventList `json:"events"`
}

type GetEventsByUserHandler struct {
	app    server.Application
	logger app.Logger
}

func (geu *GetEventsByUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.PathValue("userId")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		geu.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := getEventsByUserResponse{
		Events: geu.app.GetEventsByUserID(userID),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func NewGetEventsByUserHandler(app server.Application, logger app.Logger) *GetEventsByUserHandler {
	return &GetEventsByUserHandler{
		app:    app,
		logger: logger,
	}
}
