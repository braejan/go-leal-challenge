package usecases

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/repository"
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/repository/postgres"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessUsecases interface {
	CreateBusiness(business model.Business) (err error)
	GetBusiness(ID uuid.UUID) (business model.Business, err error)
	GetAllBusinesses() (businesses []model.Business, err error)
}

type businessUsecases struct {
	repo repository.BusinessRepository
}

func NewBusinessUsecases(db *gorm.DB) (usecases BusinessUsecases) {
	repo := postgres.NewPostgresBusinessRepository(db)
	return NewBusinessUsecasesWithRepo(repo)
}

func NewBusinessUsecasesWithRepo(repository repository.BusinessRepository) BusinessUsecases {
	return &businessUsecases{
		repo: repository,
	}
}

func (b *businessUsecases) CreateBusiness(business model.Business) (err error) {
	return b.repo.CreateBusiness(context.Background(), &business)
}

func (b *businessUsecases) GetBusiness(ID uuid.UUID) (business model.Business, err error) {
	found, err := b.repo.GetBusinessByID(context.Background(), ID)
	if err != nil {
		return
	}
	business = *found
	return
}

func (b *businessUsecases) GetAllBusinesses() (businesses []model.Business, err error) {
	found, err := b.repo.ListBusinesses(context.Background())
	if err != nil {
		return
	}
	businesses = make([]model.Business, len(found))
	for i, business := range found {
		businesses[i] = *business
	}
	return
}
