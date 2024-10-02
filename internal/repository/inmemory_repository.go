package repository

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/app/models"
	"sync"
)

type InMemoryRepo struct {
	Users       sync.Map
	Credentials sync.Map
	Texts       sync.Map
	Binaries    sync.Map
	Cards       sync.Map
}

func newInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		Users:       sync.Map{},
		Credentials: sync.Map{},
		Texts:       sync.Map{},
		Binaries:    sync.Map{},
		Cards:       sync.Map{},
	}
}

func (d *InMemoryRepo) SaveCredentials(ctx context.Context, data models.CredentialsData) error {
	select {
	case <-ctx.Done():
		return models.ErrContextTimeout
	default:
		d.Credentials.Store(data.UUID, data)
		return nil
	}
}

func (d *InMemoryRepo) LoadCredentials(ctx context.Context, id string) (models.CredentialsData, error) {
	var result models.CredentialsData
	select {
	case <-ctx.Done():
		return result, models.ErrContextTimeout
	default:
		data, ok := d.Credentials.Load(id)
		if !ok {
			return result, models.ErrInMemoryDB
		}
		result = data.(models.CredentialsData)
		return result, nil
	}
}

func (d *InMemoryRepo) Ping(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return models.ErrDatabaseUnreachable
	default:
		return nil
	}
}

func (d *InMemoryRepo) SaveUser(ctx context.Context, data models.User) error {
	select {
	case <-ctx.Done():
		return models.ErrContextTimeout
	default:
		d.Users.Store(data.UUID, data)
		return nil
	}
}

func (d *InMemoryRepo) LoadUser(ctx context.Context, id string) (models.User, error) {
	select {
	case <-ctx.Done():
		return models.User{}, models.ErrContextTimeout
	default:
		data, ok := d.Users.Load(id)
		if !ok {
			return models.User{}, models.ErrUserNotFound
		}
		return data.(models.User), nil
	}
}

func (d *InMemoryRepo) FindUser(ctx context.Context, login, password string) (models.User, error) {
	select {
	case <-ctx.Done():
		return models.User{}, models.ErrContextTimeout
	default:
		requiredUserIDs := make([]string, 0)
		d.Users.Range(func(key, value any) bool {
			if value.(models.User).Login == login && value.(models.User).Password == password {
				requiredUserIDs = append(requiredUserIDs, key.(string))
				return false
			}
			return true
		})
		if len(requiredUserIDs) == 0 {
			return models.User{}, models.ErrUserNotFound
		}
		userID := requiredUserIDs[0]
		userData, err := d.LoadUser(ctx, userID)
		if err != nil {
			return models.User{}, models.ErrInMemoryDB
		}
		return userData, nil
	}
}

func (d *InMemoryRepo) SaveText(ctx context.Context, data models.TextData) error {
	select {
	case <-ctx.Done():
		return models.ErrContextTimeout
	default:
		d.Texts.Store(data.UUID, data)
		return nil
	}
}

func (d *InMemoryRepo) LoadText(ctx context.Context, id string) (models.TextData, error) {
	var result models.TextData
	select {
	case <-ctx.Done():
		return result, models.ErrContextTimeout
	default:
		data, ok := d.Texts.Load(id)
		if !ok {
			return result, models.ErrInMemoryDB
		}
		result = data.(models.TextData)
		return result, nil
	}
}

func (d *InMemoryRepo) SaveBinary(ctx context.Context, data models.BinaryData) error {
	select {
	case <-ctx.Done():
		return models.ErrContextTimeout
	default:
		d.Binaries.Store(data.UUID, data)
		return nil
	}
}

func (d *InMemoryRepo) LoadBinary(ctx context.Context, id string) (models.BinaryData, error) {
	var result models.BinaryData
	select {
	case <-ctx.Done():
		return result, models.ErrContextTimeout
	default:
		data, ok := d.Binaries.Load(id)
		if !ok {
			return result, models.ErrInMemoryDB
		}
		result = data.(models.BinaryData)
		return result, nil
	}
}

func (d *InMemoryRepo) SaveCard(ctx context.Context, data models.BankCardData) error {
	select {
	case <-ctx.Done():
		return models.ErrContextTimeout
	default:
		d.Cards.Store(data.UUID, data)
		return nil
	}
}

func (d *InMemoryRepo) LoadCard(ctx context.Context, id string) (models.BankCardData, error) {
	var result models.BankCardData
	select {
	case <-ctx.Done():
		return result, models.ErrContextTimeout
	default:
		data, ok := d.Cards.Load(id)
		if !ok {
			return result, models.ErrInMemoryDB
		}
		result = data.(models.BankCardData)
		return result, nil
	}
}

func (d *InMemoryRepo) Add(ctx context.Context, session models.Session) error {
	return nil
}

func (d *InMemoryRepo) Remove(ctx context.Context, session models.Session) error {
	return nil
}
