package postgres

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresCampaignRepository struct {
	db *gorm.DB
}

func NewPostgresCampaignRepository(db *gorm.DB) repository.CampaignRepository {
	return &postgresCampaignRepository{
		db: db,
	}
}

func (repo *postgresCampaignRepository) CreateCampaign(ctx context.Context, campaign *model.Campaign) (err error) {
	if campaign == nil {
		err = gorm.ErrInvalidValue
		return
	}
	return repo.db.WithContext(ctx).Create(campaign).Error
}
func (repo *postgresCampaignRepository) GetCampaignByID(ctx context.Context, ID uuid.UUID) (campaign *model.Campaign, err error) {
	if ID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	campaign = &model.Campaign{
		ID: ID,
	}
	err = repo.db.WithContext(ctx).First(campaign).Error
	if err != nil {
		return
	}
	return
}
func (repo *postgresCampaignRepository) UpdateCampaign(ctx context.Context, campaign *model.Campaign) (err error) {
	_, err = repo.GetCampaignByID(ctx, campaign.ID)
	if err != nil {
		return
	}
	return repo.db.WithContext(ctx).Save(campaign).Error
}
func (repo *postgresCampaignRepository) DeleteCampaign(ctx context.Context, ID uuid.UUID) (err error) {
	_, err = repo.GetCampaignByID(ctx, ID)
	if err != nil {
		return
	}
	campaign := &model.Campaign{
		ID: ID,
	}
	return repo.db.WithContext(ctx).Delete(campaign).Error
}

func (repo *postgresCampaignRepository) ListCampaignsByBusinessID(ctx context.Context, businessID uuid.UUID) (campaigns []*model.Campaign, err error) {
	if businessID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	campaign := &model.Campaign{
		BusinessID: businessID,
	}
	err = repo.db.WithContext(ctx).Find(&campaigns, campaign).Error
	if err != nil {
		return nil, err
	}
	return
}

func (repo *postgresCampaignRepository) ListCampaignsByBranchID(ctx context.Context, branchID uuid.UUID) (campaigns []*model.Campaign, err error) {
	if branchID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	campaign := &model.Campaign{
		BranchID: branchID,
	}
	err = repo.db.WithContext(ctx).Find(&campaigns, campaign).Error
	if err != nil {
		return nil, err
	}
	return
}
