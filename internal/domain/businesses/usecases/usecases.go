package usecases

import (
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"github.com/google/uuid"
)

type BusinessUsecases interface {
	CreateBusiness(business model.Business) (ID uuid.UUID, err error)
	GetBusiness(ID uuid.UUID) (business model.Business, err error)
}
