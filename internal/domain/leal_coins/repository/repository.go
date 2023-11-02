package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/leal_coins/model"
	"github.com/google/uuid"
)

type LealCoinRepository interface {
	CreateLealCoin(ctx context.Context, lealCoin *model.LealCoin) error
	GetLealCoinByID(ctx context.Context, id uuid.UUID) (*model.LealCoin, error)
	UpdateLealCoin(ctx context.Context, lealCoin *model.LealCoin) error
	DeleteLealCoin(ctx context.Context, id uuid.UUID) error
}
