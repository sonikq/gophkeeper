package usecase

import "context"

type GophKeeperUseCase interface {
	Ping(ctx context.Context) error
}
