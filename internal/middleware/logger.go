package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type wrappedResponseWriter struct {
	http.ResponseWriter

	status int
}

func (wrw *wrappedResponseWriter) WriteHeader(statusCode int) {
	wrw.status = statusCode
	wrw.ResponseWriter.WriteHeader(statusCode)
}

func (wrw *wrappedResponseWriter) Write(data []byte) (int, error) {
	if wrw.status == 0 {
		wrw.status = http.StatusOK
	}

	return wrw.ResponseWriter.Write(data)
}

func (wrw *wrappedResponseWriter) Status() int {
	if wrw.status == 0 {
		return http.StatusOK
	}

	return wrw.status
}

func LoggerMiddleware(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrw := &wrappedResponseWriter{
				ResponseWriter: w,
				status:         0,
			}

			next.ServeHTTP(wrw, r)

			duration := time.Since(start)

			attrs := []slog.Attr{
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.Int("status", wrw.Status()),
				slog.String("duration", duration.String()),
			}

			log.LogAttrs(r.Context(), slog.LevelInfo, "HTTP request", attrs...)
		})
	}
}
