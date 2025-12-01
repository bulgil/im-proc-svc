package routes

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/bulgil/im-proc-svc/internal/http/handlers/auth"
	"github.com/bulgil/im-proc-svc/internal/repository/user"
	"github.com/bulgil/im-proc-svc/internal/validator"
)

func RegisterRoutes(srv *http.Server, userRepo *user.Repository, val *validator.Validator, log *slog.Logger) error {
	if srv.Handler == nil {
		mux := http.NewServeMux()
		registerRoutes(mux, userRepo, val, log)
		srv.Handler = mux
		return nil
	}

	mux, ok := srv.Handler.(*http.ServeMux)
	if !ok {
		return errors.New("couldn't to register routes, handler is not ServeMux")
	}

	registerRoutes(mux, userRepo, val, log)
	return nil
}

func registerRoutes(mux *http.ServeMux, userRepo *user.Repository, val *validator.Validator, log *slog.Logger) {
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})

	mux.Handle("POST /register", auth.Register(userRepo, val, log))
}
