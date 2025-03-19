package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
)

type getEventsByUserAndDatesResponse struct {
	Events storage.EventList `json:"events"`
}

type GetEventsByUserAndDatesHandler struct {
	app    server.Application
	logger app.Logger
}

func (gd *GetEventsByUserAndDatesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.PathValue("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		gd.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dateFromStr := r.PathValue("dateFrom")
	dateFrom, err := strconv.ParseInt(dateFromStr, 10, 64)
	if err != nil {
		gd.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dateToStr := r.PathValue("dateTo")
	dateTo, err := strconv.ParseInt(dateToStr, 10, 64)
	if err != nil {
		gd.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fromTime := time.Unix(dateFrom, 0)
	toTime := time.Unix(dateTo, 0)

	response := getEventsByUserAndDatesResponse{
		Events: gd.app.GetEventsByUserIDAndDates(userID, fromTime, toTime),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func NewGetEventsByUserAndDatesHandler(app server.Application, logger app.Logger) *GetEventsByUserAndDatesHandler {
	return &GetEventsByUserAndDatesHandler{
		app:    app,
		logger: logger,
	}
}
