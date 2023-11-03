package campaigns

import (
	"github.com/braejan/go-leal-challenge/internal/db/postgres"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/usecases"
)

var (
	campaignUsecases usecases.CampaignUsecases
)

func getCampaignUsecases() (c usecases.CampaignUsecases, err error) {
	if campaignUsecases == nil {
		db, err := postgres.OpenDefautDatabase()
		if err != nil {
			return nil, err
		}
		campaignUsecases = usecases.NewCampaignUsecases(db)
	}
	return campaignUsecases, nil
}
