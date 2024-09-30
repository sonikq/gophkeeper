package usecase

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/repository"
)

type gophKeeperUseCase struct {
	repo repository.Repository
}

func NewGophKeeperUseCase(r repository.Repository) GophKeeperUseCase {
	return &gophKeeperUseCase{repo: r}
}

func (g *gophKeeperUseCase) Ping(ctx context.Context) error {
	return g.repo.Ping(ctx)
}
