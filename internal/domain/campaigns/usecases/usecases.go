package usecases

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository/postgres"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignUsecases interface {
	GetCampaignByID(ID uuid.UUID) (model.Campaign, error)
	CreateNewCampaign(campaign model.Campaign) error
	GetCampaignsByBusinessID(ID uuid.UUID) ([]model.Campaign, error)
	GetCampaignsByBranchID(ID uuid.UUID) ([]model.Campaign, error)
}

type campaingUsecases struct {
	repo repository.CampaignRepository
}

func NewCampaignUsecases(db *gorm.DB) CampaignUsecases {
	repo := postgres.NewPostgresCampaignRepository(db)
	return &campaingUsecases{
		repo: repo,
	}
}

func (u *campaingUsecases) GetCampaignByID(ID uuid.UUID) (campaign model.Campaign, err error) {
	found, err := u.repo.GetCampaignByID(context.Background(), ID)
	if err != nil {
		return
	}
	campaign = *found
	return
}
func (u *campaingUsecases) CreateNewCampaign(campaign model.Campaign) (err error) {
	return u.repo.CreateCampaign(context.Background(), &campaign)
}
func (u *campaingUsecases) GetCampaignsByBusinessID(ID uuid.UUID) (campaigns []model.Campaign, err error) {
	found, err := u.repo.ListCampaignsByBusinessID(context.Background(), ID)
	if err != nil {
		return
	}
	campaigns = make([]model.Campaign, len(found))
	for i, c := range found {
		campaigns[i] = *c
	}
	return
}
func (u *campaingUsecases) GetCampaignsByBranchID(ID uuid.UUID) (campaigns []model.Campaign, err error) {
	found, err := u.repo.ListCampaignsByBranchID(context.Background(), ID)
	if err != nil {
		return
	}
	campaigns = make([]model.Campaign, len(found))
	for i, c := range found {
		campaigns[i] = *c
	}
	return
}
