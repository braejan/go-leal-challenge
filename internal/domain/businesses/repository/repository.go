package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"github.com/google/uuid"
)

type BusinessRepository interface {
	CreateBusiness(ctx context.Context, business *model.Business) error
	GetBusinessByID(ctx context.Context, id uuid.UUID) (*model.Business, error)
	UpdateBusiness(ctx context.Context, business *model.Business) error
	DeleteBusiness(ctx context.Context, id uuid.UUID) error
	ListBusinesses(ctx context.Context) ([]*model.Business, error)
}
