package repository

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/app/server/config"
)

func NewGophKeeperRepository(ctx context.Context, cfg config.PostgresConfig) Repository {
	return NewPostgresRepository(ctx, cfg) // realise with options pattern
}
