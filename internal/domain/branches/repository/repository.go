package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
)

type BranchRepository interface {
	CreateBranch(ctx context.Context, branch *model.Branch) (int, error)
	GetBranchByID(ctx context.Context, branchID int) (*model.Branch, error)
	GetBranchesByComercioID(ctx context.Context, comercioID int) ([]*model.Branch, error)
	UpdateBranch(ctx context.Context, branch *model.Branch) error
	DeleteBranch(ctx context.Context, branchID int) error
}
