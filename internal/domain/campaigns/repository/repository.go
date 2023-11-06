package repository

import (
	"context"
	"time"

	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/google/uuid"
)

type CampaignRepository interface {
	CreateCampaign(ctx context.Context, campaign *model.Campaign) error
	GetCampaignByID(ctx context.Context, id uuid.UUID) (*model.Campaign, error)
	UpdateCampaign(ctx context.Context, campaign *model.Campaign) error
	DeleteCampaignByID(ctx context.Context, id uuid.UUID) error
	ListCampaignsByBusinessID(ctx context.Context, businessID uuid.UUID) ([]*model.Campaign, error)
	ListCampaignsByBranchID(ctx context.Context, branchID uuid.UUID) ([]*model.Campaign, error)
	GetUnfinishedCampaigns(ctx context.Context, txDate time.Time, branchID uuid.UUID) (campaigns []*model.Campaign, err error)
}
