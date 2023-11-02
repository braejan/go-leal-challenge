package postgres

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/braejan/go-leal-challenge/internal/domain/branches/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresBranchRepository struct {
	db *gorm.DB
}

func NewPostgresBranchRepository(db *gorm.DB) repository.BranchRepository {
	return &postgresBranchRepository{db}
}

func (repo *postgresBranchRepository) CreateBranch(ctx context.Context, branch *model.Branch) (err error) {
	if branch == nil {
		err = gorm.ErrInvalidValue
		return
	}
	return repo.db.WithContext(ctx).Create(branch).Error
}
func (repo *postgresBranchRepository) GetBranchByID(ctx context.Context, ID uuid.UUID) (branch *model.Branch, err error) {
	if ID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	branch = &model.Branch{
		ID: ID,
	}
	err = repo.db.WithContext(ctx).First(branch).Error
	if err != nil {
		branch = nil
	}
	return
}
func (repo *postgresBranchRepository) GetBranchesByBusinessID(ctx context.Context, businessID uuid.UUID) (branches []*model.Branch, err error) {
	if businessID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	branch := &model.Branch{
		BusinessID: businessID,
	}
	err = repo.db.WithContext(ctx).Find(&branches, branch).Error
	if err != nil {
		branches = nil
	}
	return
}
func (repo *postgresBranchRepository) UpdateBranch(ctx context.Context, branch *model.Branch) (err error) {
	if branch == nil {
		err = gorm.ErrInvalidValue
		return
	}
	_, err = repo.GetBranchByID(ctx, branch.ID)
	if err != nil {
		return
	}
	return repo.db.WithContext(ctx).Save(branch).Error
}
func (repo *postgresBranchRepository) DeleteBranchByID(ctx context.Context, ID uuid.UUID) (err error) {
	if ID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	_, err = repo.GetBranchByID(ctx, ID)
	if err != nil {
		return
	}
	return repo.db.WithContext(ctx).Delete(&model.Branch{ID: ID}).Error
}
