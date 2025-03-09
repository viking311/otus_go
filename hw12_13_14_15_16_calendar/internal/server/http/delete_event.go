package internalhttp

import (
	"net/http"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type DeleteEventHandler struct {
	app    Application
	logger app.Logger
}

func (de *DeleteEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	de.app.DeleteEvent(id)
}

func NewDeleteEventHandler(app Application, logger app.Logger) *DeleteEventHandler {
	return &DeleteEventHandler{
		app:    app,
		logger: logger,
	}
}
