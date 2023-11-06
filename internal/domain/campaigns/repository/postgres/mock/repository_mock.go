package mock

import (
	"context"
	"time"

	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// type CampaignRepository interface {
// 	CreateCampaign(ctx context.Context, campaign *model.Campaign) error
// 	GetCampaignByID(ctx context.Context, id uuid.UUID) (*model.Campaign, error)
// 	UpdateCampaign(ctx context.Context, campaign *model.Campaign) error
// 	DeleteCampaignByID(ctx context.Context, id uuid.UUID) error
// 	ListCampaignsByBusinessID(ctx context.Context, businessID uuid.UUID) ([]*model.Campaign, error)
// 	ListCampaignsByBranchID(ctx context.Context, branchID uuid.UUID) ([]*model.Campaign, error)
// 	GetUncompletedCampaigns(ctx context.Context) (campaigns []*model.Campaign, err error)
// }

type mockCampaignRepository struct {
	mock.Mock
}

func NewMockCampaignRepository() *mockCampaignRepository {
	return &mockCampaignRepository{}
}

func (m *mockCampaignRepository) CreateCampaign(ctx context.Context, campaign *model.Campaign) error {
	args := m.Called(ctx, campaign)
	return args.Error(0)
}
func (m *mockCampaignRepository) GetCampaignByID(ctx context.Context, id uuid.UUID) (*model.Campaign, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Campaign), args.Error(1)
}
func (m *mockCampaignRepository) UpdateCampaign(ctx context.Context, campaign *model.Campaign) error {
	args := m.Called(ctx, campaign)
	return args.Error(0)
}
func (m *mockCampaignRepository) DeleteCampaignByID(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *mockCampaignRepository) ListCampaignsByBusinessID(ctx context.Context, businessID uuid.UUID) ([]*model.Campaign, error) {
	args := m.Called(ctx, businessID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Campaign), args.Error(1)
}
func (m *mockCampaignRepository) ListCampaignsByBranchID(ctx context.Context, branchID uuid.UUID) ([]*model.Campaign, error) {
	args := m.Called(ctx, branchID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Campaign), args.Error(1)
}
func (m *mockCampaignRepository) GetUnfinishedCampaigns(ctx context.Context, txDate time.Time, branchID uuid.UUID) ([]*model.Campaign, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Campaign), args.Error(1)
}
