package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/google/uuid"
)

type BranchRepository interface {
	CreateBranch(ctx context.Context, branch *model.Branch) error
	GetBranchByID(ctx context.Context, ID uuid.UUID) (*model.Branch, error)
	GetBranchesByBusinessID(ctx context.Context, businessID uuid.UUID) ([]*model.Branch, error)
	UpdateBranch(ctx context.Context, branch *model.Branch) error
	DeleteBranchByID(ctx context.Context, ID uuid.UUID) error
}
