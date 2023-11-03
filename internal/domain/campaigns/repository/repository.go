package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/google/uuid"
)

type CampaignRepository interface {
	CreateCampaign(ctx context.Context, campaign *model.Campaign) error
	GetCampaignByID(ctx context.Context, id uuid.UUID) (*model.Campaign, error)
	UpdateCampaign(ctx context.Context, campaign *model.Campaign) error
	DeleteCampaign(ctx context.Context, id uuid.UUID) error
	ListCampaignsByBusinessID(ctx context.Context, businessID uuid.UUID) ([]*model.Campaign, error)
	ListCampaignsByBranchID(ctx context.Context, branchID uuid.UUID) ([]*model.Campaign, error)
	GetUncompletedCampaigns(ctx context.Context) (campaigns []*model.Campaign, err error)
}
