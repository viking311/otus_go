package internalhttp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type LoggerMiddleware struct {
	logger app.Logger
}

func (lm *LoggerMiddleware) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wr := NewWrapResponseWriter(w)
		next.ServeHTTP(wr, r)
		msg := fmt.Sprintf("%s %s %s %s %d %s %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			r.Proto,
			wr.statusCode,
			time.Since(start).String(),
			r.UserAgent())
		lm.logger.Info(msg)
	})
}

func NewLoggingMiddleware(logger app.Logger) LoggerMiddleware {
	return LoggerMiddleware{
		logger: logger,
	}
}
