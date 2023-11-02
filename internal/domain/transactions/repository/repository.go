package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/leal_points/model"
	"github.com/google/uuid"
)

type LealPointRepository interface {
	CreateLealPoint(ctx context.Context, lealPoint *model.LealPoint) error
	GetLealPointByID(ctx context.Context, id uuid.UUID) (*model.LealPoint, error)
	UpdateLealPoint(ctx context.Context, lealPoint *model.LealPoint) error
	DeleteLealPoint(ctx context.Context, id uuid.UUID) error
}
