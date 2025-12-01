package application

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/bulgil/im-proc-svc/internal/config"
	"github.com/bulgil/im-proc-svc/internal/server"
	"github.com/bulgil/im-proc-svc/internal/validator"
)

type App struct {
	Logger    *slog.Logger
	Server    *server.Server
	Validator *validator.Validator

	wg *sync.WaitGroup
}

func Start() {
	cfg := config.ReadConfig()
	log := newLogger(cfg.Env)

	app := &App{
		Logger:    log,
		Server:    server.New(cfg.HTTPServerCfg, log),
		Validator: validator.New(),
		wg:        &sync.WaitGroup{},
	}

	app.Logger.Info("application is starting", "env", cfg.Env)
	app.run()
}

func (a *App) run() {
	a.wg.Go(a.Server.Run)

	c := sigtermChan()
	<-c

	a.Server.Stop()

	a.wg.Wait()

	a.Logger.Info("application is stopped")
}

func newLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "dev":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case "prod":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	default:
		panic(fmt.Sprintf("%s env is not implemented", env))
	}

	return log
}

func sigtermChan() chan os.Signal {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	return c
}
