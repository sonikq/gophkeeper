package usecase

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/app/models"
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

func (g *gophKeeperUseCase) FindUser(ctx context.Context, login, password string) (models.User, error) {
	return g.repo.FindUser(ctx, login, password)
}

func (g *gophKeeperUseCase) SaveCredentials(ctx context.Context, data models.CredentialsData) error {
	return g.repo.SaveCredentials(ctx, data)
}

func (g *gophKeeperUseCase) LoadCredentials(ctx context.Context, id string) (models.CredentialsData, error) {
	return g.repo.LoadCredentials(ctx, id)
}

func (g *gophKeeperUseCase) SaveUser(ctx context.Context, data models.User) error {
	return g.repo.SaveUser(ctx, data)
}

func (g *gophKeeperUseCase) LoadUser(ctx context.Context, id string) (models.User, error) {
	return g.repo.LoadUser(ctx, id)
}

func (g *gophKeeperUseCase) SaveBinary(ctx context.Context, data models.BinaryData) error {
	return g.repo.SaveBinary(ctx, data)
}

func (g *gophKeeperUseCase) LoadBinary(ctx context.Context, id string) (models.BinaryData, error) {
	return g.repo.LoadBinary(ctx, id)
}

func (g *gophKeeperUseCase) SaveText(ctx context.Context, data models.TextData) error {
	return g.repo.SaveText(ctx, data)
}

func (g *gophKeeperUseCase) LoadText(ctx context.Context, id string) (models.TextData, error) {
	return g.repo.LoadText(ctx, id)
}

func (g *gophKeeperUseCase) SaveCard(ctx context.Context, data models.BankCardData) error {
	return g.repo.SaveCard(ctx, data)
}

func (g *gophKeeperUseCase) LoadCard(ctx context.Context, id string) (models.BankCardData, error) {
	return g.repo.LoadCard(ctx, id)
}
