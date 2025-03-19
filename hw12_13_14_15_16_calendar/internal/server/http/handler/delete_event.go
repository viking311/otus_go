package handler

import (
	"net/http"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"
)

type DeleteEventHandler struct {
	app    server.Application
	logger app.Logger
}

func (de *DeleteEventHandler) ServeHTTP(_ http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	de.app.DeleteEvent(id)
}

func NewDeleteEventHandler(app server.Application, logger app.Logger) *DeleteEventHandler {
	return &DeleteEventHandler{
		app:    app,
		logger: logger,
	}
}
