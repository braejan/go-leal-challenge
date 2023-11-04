package usecases

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/braejan/go-leal-challenge/internal/domain/branches/repository"
	"github.com/braejan/go-leal-challenge/internal/domain/branches/repository/postgres"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BranchUsecases interface {
	CreateBranch(branch model.Branch) (err error)
	GetBranchByID(ID uuid.UUID) (branch model.Branch, err error)
	GetBranchesByBusinessID(ID uuid.UUID) (branches []model.Branch, err error)
}

type branchUsecases struct {
	repo repository.BranchRepository
}

func NewBranchUsecases(db *gorm.DB) BranchUsecases {
	repo := postgres.NewPostgresBranchRepository(db)
	return NewBranchUsecasesWithRepo(repo)
}

func NewBranchUsecasesWithRepo(repository repository.BranchRepository) BranchUsecases {
	return &branchUsecases{
		repo: repository,
	}
}

func (u *branchUsecases) CreateBranch(branch model.Branch) (err error) {
	return u.repo.CreateBranch(context.Background(), &branch)
}

func (u *branchUsecases) GetBranchByID(ID uuid.UUID) (branch model.Branch, err error) {
	found, err := u.repo.GetBranchByID(context.Background(), ID)
	if err != nil {
		return
	}
	branch = *found
	return
}

func (u *branchUsecases) GetBranchesByBusinessID(ID uuid.UUID) (branches []model.Branch, err error) {
	found, err := u.repo.GetBranchesByBusinessID(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	branches = make([]model.Branch, len(found))
	for i, branch := range found {
		branches[i] = *branch
	}
	return
}
