package util

import (
	modelCampaign "github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	modelTransaction "github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
)

func HasCampaign(tx *modelTransaction.Transaction, campaign *modelCampaign.Campaign) (apply bool) {
	if campaign.Completed {
		return
	}
	if tx.Amount < campaign.MinAmount {
		return
	}
	if tx.Accumulated || tx.Redeemed {
		return
	}
	if tx.PurchaseDate.Before(campaign.StartDate) || tx.PurchaseDate.After(campaign.EndDate) {
		return
	}
	return true
}

func CalcPointsAndCashbackAccumulate(tx *modelTransaction.Transaction, campaign *modelCampaign.Campaign, conversionFactor float64, equivalence float64) (points int, cashback float64) {
	if !HasCampaign(tx, campaign) {
		return
	}
	points = int(tx.Amount * conversionFactor * campaign.RewardMultiplier)
	cashback = tx.Amount * equivalence * campaign.RewardMultiplier
	return
}
