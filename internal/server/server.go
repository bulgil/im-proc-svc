package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/bulgil/im-proc-svc/internal/config"
	"github.com/bulgil/im-proc-svc/internal/middleware"
	"github.com/bulgil/im-proc-svc/internal/routes"
)

type Server struct {
	httpServer *http.Server
	log        *slog.Logger
}

func New(httpCfg config.HTTPServerCfg, log *slog.Logger) *Server {
	srv := &http.Server{
		Addr: httpCfg.Addr,
	}

	routes.RegisterRoutes(srv)
	middleware.ApplyMiddlewares(srv, log)

	return &Server{
		httpServer: srv,
		log:        log,
	}
}

func (s *Server) Run() {
	s.log.Info("HTTP server started", "addr", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.Error("HTTP server failed", "error", err.Error())
		return
	}
}

func (s *Server) Stop() {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := s.httpServer.Shutdown(shutdownCtx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			s.log.Error("server shutdown timeout exceeded: some connections may not be closed properly")
			return
		}

		s.log.Error("server shutted down with a problem", "error", err.Error())
		return
	}

	s.log.Info("server is gracefully shutdowned")
}
