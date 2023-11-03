package usecases

import (
	"context"

	branchRepo "github.com/braejan/go-leal-challenge/internal/domain/branches/repository"
	branchPostgres "github.com/braejan/go-leal-challenge/internal/domain/branches/repository/postgres"
	businessRepo "github.com/braejan/go-leal-challenge/internal/domain/businesses/repository"
	businessPostgres "github.com/braejan/go-leal-challenge/internal/domain/businesses/repository/postgres"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	campaignRepo "github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository"
	campaignPostgres "github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository/postgres"
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
	campaignRepo.CampaignRepository
	businessRepo.BusinessRepository
	branchRepo.BranchRepository
}

func NewCampaignUsecases(db *gorm.DB) CampaignUsecases {
	repoCampaign := campaignPostgres.NewPostgresCampaignRepository(db)
	repoBusiness := businessPostgres.NewPostgresBusinessRepository(db)
	repoBranch := branchPostgres.NewPostgresBranchRepository(db)
	return &campaingUsecases{
		CampaignRepository: repoCampaign,
		BusinessRepository: repoBusiness,
		BranchRepository:   repoBranch,
	}
}

func (u *campaingUsecases) GetCampaignByID(ID uuid.UUID) (campaign model.Campaign, err error) {
	found, err := u.CampaignRepository.GetCampaignByID(context.Background(), ID)
	if err != nil {
		return
	}
	campaign = *found
	return
}
func (u *campaingUsecases) CreateNewCampaign(campaign model.Campaign) (err error) {
	found, err := u.BranchRepository.GetBranchByID(context.Background(), *campaign.BranchID)
	if err != nil {
		return
	}
	if *found.BusinessID != *campaign.BusinessID {
		err = gorm.ErrInvalidValue
		return
	}
	return u.CampaignRepository.CreateCampaign(context.Background(), &campaign)
}
func (u *campaingUsecases) GetCampaignsByBusinessID(ID uuid.UUID) (campaigns []model.Campaign, err error) {
	_, err = u.BusinessRepository.GetBusinessByID(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	found, err := u.CampaignRepository.ListCampaignsByBusinessID(context.Background(), ID)
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
	_, err = u.BranchRepository.GetBranchByID(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	found, err := u.CampaignRepository.ListCampaignsByBranchID(context.Background(), ID)
	if err != nil {
		return
	}
	campaigns = make([]model.Campaign, len(found))
	for i, c := range found {
		campaigns[i] = *c
	}
	return
}
