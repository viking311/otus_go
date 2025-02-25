package internalhttp

import "net/http"

type WrapResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *WrapResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func NewWrapResponseWriter(w http.ResponseWriter) *WrapResponseWriter {
	return &WrapResponseWriter{w, http.StatusOK}
}
