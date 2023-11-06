package postgres

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresBusinessRepository struct {
	db *gorm.DB
}

func NewPostgresBusinessRepository(db *gorm.DB) repository.BusinessRepository {
	return &postgresBusinessRepository{db: db}
}

func (repo *postgresBusinessRepository) CreateBusiness(ctx context.Context, business *model.Business) (err error) {
	if business == nil {
		err = gorm.ErrInvalidValue
		return
	}
	return repo.db.WithContext(ctx).Create(business).Error
}
func (repo *postgresBusinessRepository) GetBusinessByID(ctx context.Context, ID uuid.UUID) (business *model.Business, err error) {
	if ID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	business = &model.Business{
		ID: &ID,
	}
	err = repo.db.WithContext(ctx).First(business).Error
	if err != nil {
		business = nil
	}
	return
}
func (repo *postgresBusinessRepository) UpdateBusiness(ctx context.Context, business *model.Business) (err error) {
	if business == nil {
		err = gorm.ErrInvalidValue
		return
	}
	_, err = repo.GetBusinessByID(ctx, *business.ID)
	if err != nil {
		return
	}
	return repo.db.WithContext(ctx).Save(business).Error
}
func (repo *postgresBusinessRepository) DeleteBusinessByID(ctx context.Context, ID uuid.UUID) (err error) {
	_, err = repo.GetBusinessByID(ctx, ID)
	if err != nil {
		return
	}
	business := &model.Business{
		ID: &ID,
	}
	return repo.db.WithContext(ctx).Delete(business).Error
}
func (repo *postgresBusinessRepository) ListBusinesses(ctx context.Context) (businesses []*model.Business, err error) {
	err = repo.db.WithContext(ctx).Find(&businesses).Error
	if err != nil {
		businesses = nil
	}
	return
}
