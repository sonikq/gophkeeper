package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonikq/gophkeeper/internal/app/server/config"
	"log"
	"time"
)

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(ctx context.Context, cfg config.PostgresConfig) Repository {
	pool, err := newRepository(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	return &postgresRepository{pool: pool}
}

func newRepository(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	t1 := time.Now()
	var pool *pgxpool.Pool
	connection, err := pgxpool.ParseConfig(cfg.URI)
	if err != nil {
		return nil, err
	}
	connection.MaxConns = int32(cfg.Workers)

	pool, err = pgxpool.NewWithConfig(ctx, connection)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Printf("connection to database took: %v\n", time.Since(t1))

	return pool, nil

}

func (r *postgresRepository) Ping(ctx context.Context) error {
	return r.pool.Ping(ctx)
}

func (r *postgresRepository) Close() {
	r.pool.Close()
}
