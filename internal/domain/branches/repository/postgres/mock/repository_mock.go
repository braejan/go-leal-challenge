package mock

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// CreateBranch(ctx context.Context, branch *model.Branch) error
// 	GetBranchByID(ctx context.Context, ID uuid.UUID) (*model.Branch, error)
// 	GetBranchesByBusinessID(ctx context.Context, businessID uuid.UUID) ([]*model.Branch, error)
// 	UpdateBranch(ctx context.Context, branch *model.Branch) error
// 	DeleteBranchByID(ctx context.Context, ID uuid.UUID) error

type mockBranchRepository struct {
	mock.Mock
}

func NewMockBranchRepository() *mockBranchRepository {
	return &mockBranchRepository{}
}

func (m *mockBranchRepository) CreateBranch(ctx context.Context, branch *model.Branch) error {
	args := m.Called(ctx, branch)
	return args.Error(0)
}

func (m *mockBranchRepository) GetBranchByID(ctx context.Context, ID uuid.UUID) (*model.Branch, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Branch), args.Error(1)
}

func (m *mockBranchRepository) GetBranchesByBusinessID(ctx context.Context, businessID uuid.UUID) ([]*model.Branch, error) {
	args := m.Called(ctx, businessID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Branch), args.Error(1)
}

func (m *mockBranchRepository) UpdateBranch(ctx context.Context, branch *model.Branch) error {
	args := m.Called(ctx, branch)
	return args.Error(0)
}

func (m *mockBranchRepository) DeleteBranchByID(ctx context.Context, ID uuid.UUID) error {
	args := m.Called(ctx, ID)
	return args.Error(0)
}
