package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is the shared PostgreSQL connection pool. Use it anywhere: internal.DB
var DB *pgxpool.Pool

// PgConn opens the connection pool and checks the database is reachable.
func PgConn(ctx context.Context, url string) error {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return fmt.Errorf("postgres: parse url: %w", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return fmt.Errorf("postgres: create pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return fmt.Errorf("postgres: ping: %w", err)
	}

	DB = pool
	return nil
}
