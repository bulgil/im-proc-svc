package routes

import (
	"errors"
	"net/http"
)

func RegisterRoutes(srv *http.Server) error {
	if srv.Handler == nil {
		mux := http.NewServeMux()
		registerRoutes(mux)
		srv.Handler = mux
		return nil
	}

	mux, ok := srv.Handler.(*http.ServeMux)
	if !ok {
		return errors.New("couldn't to register routes, handler is not ServeMux")
	}

	registerRoutes(mux)
	return nil
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})
}
