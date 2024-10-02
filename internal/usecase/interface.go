package usecase

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/app/models"
)

type GophKeeperUseCase interface {
	Ping(ctx context.Context) error
	FindUser(ctx context.Context, login, password string) (models.User, error)
	Credentials
	Auth
	Binaries
	Texts
	Cards
}

type Credentials interface {
	SaveCredentials(ctx context.Context, data models.CredentialsData) error
	LoadCredentials(ctx context.Context, id string) (models.CredentialsData, error)
}

type Auth interface {
	SaveUser(ctx context.Context, data models.User) error
	LoadUser(ctx context.Context, id string) (models.User, error)
}

type Binaries interface {
	SaveBinary(ctx context.Context, data models.BinaryData) error
	LoadBinary(ctx context.Context, id string) (models.BinaryData, error)
}

type Texts interface {
	SaveText(ctx context.Context, data models.TextData) error
	LoadText(ctx context.Context, id string) (models.TextData, error)
}

type Cards interface {
	SaveCard(ctx context.Context, data models.BankCardData) error
	LoadCard(ctx context.Context, id string) (models.BankCardData, error)
}
