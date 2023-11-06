package mock

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// type BusinessRepository interface {
// 	CreateBusiness(ctx context.Context, business *model.Business) error
// 	GetBusinessByID(ctx context.Context, ID uuid.UUID) (*model.Business, error)
// 	UpdateBusiness(ctx context.Context, business *model.Business) error
// 	DeleteBusinessByID(ctx context.Context, ID uuid.UUID) error
// 	ListBusinesses(ctx context.Context) ([]*model.Business, error)
// }

type mockBusinessRepository struct {
	mock.Mock
}

func NewMockBusinessRepository() *mockBusinessRepository {
	return &mockBusinessRepository{}
}

func (m *mockBusinessRepository) CreateBusiness(ctx context.Context, business *model.Business) error {
	args := m.Called(ctx, business)
	return args.Error(0)
}

func (m *mockBusinessRepository) GetBusinessByID(ctx context.Context, ID uuid.UUID) (*model.Business, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Business), args.Error(1)
}

func (m *mockBusinessRepository) UpdateBusiness(ctx context.Context, business *model.Business) error {
	args := m.Called(ctx, business)
	return args.Error(0)
}

func (m *mockBusinessRepository) DeleteBusinessByID(ctx context.Context, ID uuid.UUID) error {
	args := m.Called(ctx, ID)
	return args.Error(0)
}

func (m *mockBusinessRepository) ListBusinesses(ctx context.Context) ([]*model.Business, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Business), args.Error(1)
}
