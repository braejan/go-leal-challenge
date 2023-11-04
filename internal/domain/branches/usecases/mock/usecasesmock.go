package mock

import (
	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type mockBranchUsecases struct {
	mock.Mock
}

// type BranchUsecases interface {
// 	CreateBranch(branch model.Branch) (err error)
// 	GetBranchByID(ID uuid.UUID) (branch model.Branch, err error)
// 	GetBranchesByBusinessID(ID uuid.UUID) (branches []model.Branch, err error)
// }

func (m *mockBranchUsecases) CreateBranch(branch model.Branch) (err error) {
	args := m.Called(branch)
	return args.Error(0)
}

func (m *mockBranchUsecases) GetBranchByID(ID uuid.UUID) (branch model.Branch, err error) {
	args := m.Called(ID)
	return args.Get(0).(model.Branch), args.Error(1)
}

func (m *mockBranchUsecases) GetBranchesByBusinessID(ID uuid.UUID) (branches []model.Branch, err error) {
	args := m.Called(ID)
	return args.Get(0).([]model.Branch), args.Error(1)
}
