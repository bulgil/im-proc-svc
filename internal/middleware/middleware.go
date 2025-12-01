package middleware

import (
	"log/slog"
	"net/http"
)

func ApplyMiddlewares(srv *http.Server, log *slog.Logger) error {
	if srv.Handler == nil {
		mux := http.NewServeMux()
		handler := applyMiddlewares(mux, log)
		srv.Handler = handler
		return nil
	}

	handler := applyMiddlewares(srv.Handler, log)
	srv.Handler = handler
	return nil
}

func applyMiddlewares(handler http.Handler, log *slog.Logger) http.Handler {
	handler = LoggerMiddleware(log)(handler)

	return handler
}
