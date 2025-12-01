package postgres

import (
	"context"
	"fmt"

	"github.com/bulgil/im-proc-svc/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(pgCfg config.PGCfg) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pgCfg.User,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.Database))
	if err != nil {
		panic(fmt.Sprintf("problem with postgres connection: %s", err.Error()))
	}

	return db
}
