package usecases_test

import (
	"context"
	"testing"

	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/repository/postgres/mock"
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_CreateBusiness_Err(t *testing.T) {
	repo := mock.NewMockBusinessRepository()
	usecases := usecases.NewBusinessUsecasesWithRepo(repo)
	repo.On("CreateBusiness", testifymock.Anything, testifymock.Anything).Return(gorm.ErrInvalidDB)
	err := usecases.CreateBusiness(model.Business{})
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_CreateBusiness_Success(t *testing.T) {
	repo := mock.NewMockBusinessRepository()
	usecases := usecases.NewBusinessUsecasesWithRepo(repo)
	repo.On("CreateBusiness", testifymock.Anything, testifymock.Anything).Return(nil)
	err := usecases.CreateBusiness(model.Business{})
	assert.Nil(t, err)
}

func Test_GetBusiness_Err(t *testing.T) {
	repo := mock.NewMockBusinessRepository()
	usecases := usecases.NewBusinessUsecasesWithRepo(repo)
	ID := uuid.New()
	repo.On("GetBusinessByID", context.Background(), ID).Return(nil, gorm.ErrInvalidDB)
	_, err := usecases.GetBusiness(ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_GetGetBusiness_Sucess(t *testing.T) {
	repo := mock.NewMockBusinessRepository()
	usecases := usecases.NewBusinessUsecasesWithRepo(repo)
	ID := uuid.New()
	business := &model.Business{
		ID: &ID,
	}
	repo.On("GetBusinessByID", context.Background(), ID).Return(business, nil)
	found, err := usecases.GetBusiness(ID)
	assert.Nil(t, err)
	assert.Equal(t, business.ID.String(), found.ID.String())
}

func Test_GetGetAllBusinesses_Error(t *testing.T) {
	repo := mock.NewMockBusinessRepository()
	usecases := usecases.NewBusinessUsecasesWithRepo(repo)
	repo.On("ListBusinesses", context.Background()).Return(nil, gorm.ErrInvalidDB)
	_, err := usecases.GetAllBusinesses()
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_GetAllBusiness_Sucess(t *testing.T) {
	repo := mock.NewMockBusinessRepository()
	usecases := usecases.NewBusinessUsecasesWithRepo(repo)
	ID := uuid.New()
	businesses := []*model.Business{
		{
			ID: &ID,
		},
		{
			ID: &ID,
		},
	}
	repo.On("ListBusinesses", context.Background()).Return(businesses, nil)
	found, err := usecases.GetAllBusinesses()
	assert.Nil(t, err)
	assert.Len(t, found, 2)
	assert.Equal(t, found[0].ID.String(), ID.String())
}
