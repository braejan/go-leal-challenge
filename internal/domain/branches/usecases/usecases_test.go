package usecases_test

import (
	"context"
	"testing"

	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/braejan/go-leal-challenge/internal/domain/branches/repository/postgres/mock"
	"github.com/braejan/go-leal-challenge/internal/domain/branches/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_CreateBranch_Err(t *testing.T) {
	repo := mock.NewMockBranchRepository()
	usecases := usecases.NewBranchUsecasesWithRepo(repo)
	repo.On("CreateBranch", testifymock.Anything, testifymock.Anything).Return(gorm.ErrInvalidDB)
	err := usecases.CreateBranch(model.Branch{})
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_CreateBranch_Success(t *testing.T) {
	repo := mock.NewMockBranchRepository()
	usecases := usecases.NewBranchUsecasesWithRepo(repo)
	repo.On("CreateBranch", testifymock.Anything, testifymock.Anything).Return(nil)
	err := usecases.CreateBranch(model.Branch{})
	assert.Nil(t, err)
}

func Test_GetBranchByID_Err(t *testing.T) {
	repo := mock.NewMockBranchRepository()
	usecases := usecases.NewBranchUsecasesWithRepo(repo)
	ID := uuid.New()
	repo.On("GetBranchByID", context.Background(), ID).Return(nil, gorm.ErrInvalidDB)
	_, err := usecases.GetBranchByID(ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_GetBranchByID_Sucess(t *testing.T) {
	repo := mock.NewMockBranchRepository()
	usecases := usecases.NewBranchUsecasesWithRepo(repo)
	ID := uuid.New()
	branch := &model.Branch{
		ID: &ID,
	}
	repo.On("GetBranchByID", context.Background(), ID).Return(branch, nil)
	found, err := usecases.GetBranchByID(ID)
	assert.Nil(t, err)
	assert.Equal(t, branch.ID.String(), found.ID.String())
}

func Test_GetBranchesByBusinessID_Error(t *testing.T) {
	repo := mock.NewMockBranchRepository()
	usecases := usecases.NewBranchUsecasesWithRepo(repo)
	ID := uuid.New()
	repo.On("GetBranchesByBusinessID", context.Background(), ID).Return(nil, gorm.ErrInvalidDB)
	_, err := usecases.GetBranchesByBusinessID(ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_GetBranchesByBusinessID_Sucess(t *testing.T) {
	repo := mock.NewMockBranchRepository()
	usecases := usecases.NewBranchUsecasesWithRepo(repo)
	ID := uuid.New()
	branch := &model.Branch{
		ID: &ID,
	}
	repo.On("GetBranchesByBusinessID", context.Background(), ID).Return([]*model.Branch{branch}, nil)
	found, err := usecases.GetBranchesByBusinessID(ID)
	assert.Nil(t, err)
	assert.Len(t, found, 1)
	assert.Equal(t, found[0].ID.String(), ID.String())
}
