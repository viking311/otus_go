package handler

import (
	"net/http"
)

type Stub struct{}

func (s *Stub) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
