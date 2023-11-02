package businesses

import (
	"github.com/braejan/go-leal-challenge/internal/db/postgres"
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/usecases"
)

var (
	businessUsecases usecases.BusinessUsecases
)

func getBusinessUsecases() (c usecases.BusinessUsecases, err error) {
	if businessUsecases == nil {
		db, err := postgres.OpenDefautDatabase()
		if err != nil {
			return nil, err
		}
		businessUsecases = usecases.NewBusinessUsecases(db)
	}
	return businessUsecases, nil
}
