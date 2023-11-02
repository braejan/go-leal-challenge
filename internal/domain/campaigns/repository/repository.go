package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/google/uuid"
)

type CampaignRepository interface {
	CreateCampaign(ctx context.Context, campaign model.Campaign) error
	GetCampaignByID(ctx context.Context, id uuid.UUID) (*model.Campaign, error)
	UpdateCampaign(ctx context.Context, campaign model.Campaign) error
	DeleteCampaign(ctx context.Context, id uuid.UUID) error
	ListCampaigns(ctx context.Context) ([]*model.Campaign, error)
}
