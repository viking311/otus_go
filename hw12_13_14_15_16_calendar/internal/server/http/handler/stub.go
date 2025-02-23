package handler

import (
	"fmt"
	"net/http"
)

type Stub struct{}

func (s *Stub) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "hello\n")
}
