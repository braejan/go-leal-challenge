package usecases

import (
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/google/uuid"
)

type CampaignUsecases interface {
	GetCampaignByID(ID uuid.UUID) (model.Campaign, error)
	CreateNewCampaign(campaign model.Campaign) (uuid.UUID, error)
	GetCampaignsByBusinessID(ID uuid.UUID) ([]model.Campaign, error)
	GetCampaignsByBranchID(ID uuid.UUID) ([]model.Campaign, error)
}
