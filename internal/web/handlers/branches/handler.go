package branches

import (
	"github.com/braejan/go-leal-challenge/internal/db/postgres"
	"github.com/braejan/go-leal-challenge/internal/domain/branches/usecases"
)

var (
	branchUsecases usecases.BranchUsecases
)

func getBranchUsecases() (c usecases.BranchUsecases, err error) {
	if branchUsecases == nil {
		db, err := postgres.OpenDefautDatabase()
		if err != nil {
			return nil, err
		}
		branchUsecases = usecases.NewBranchUsecases(db)
	}
	return branchUsecases, nil
}
