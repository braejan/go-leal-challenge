package usecases

import (
	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/google/uuid"
)

type BranchUsecases interface {
	CreateBranch(branch model.Branch) (ID uuid.UUID, err error)
	GetBranches() (branches []model.Branch, err error)
}
