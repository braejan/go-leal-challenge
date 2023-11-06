package util_test

import (
	"fmt"
	"testing"
	"time"

	modelCampaign "github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/util"
	modelTransaction "github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func dummyTransactions() []*modelTransaction.Transaction {
	branchID := uuid.New()
	userID := uuid.New()
	date, _ := time.Parse("2006-01-02 15:04", "2023-05-01 00:00")
	return []*modelTransaction.Transaction{
		{
			BranchID:     &branchID,
			Amount:       5000,
			UserID:       &userID,
			PurchaseDate: date.AddDate(0, 0, 1),
		},
		{
			BranchID:     &branchID,
			Amount:       23000,
			UserID:       &userID,
			PurchaseDate: date.AddDate(0, 0, 15),
		},
		{
			BranchID:     &branchID,
			Amount:       25000,
			UserID:       &userID,
			PurchaseDate: date.AddDate(0, 0, 25),
		},
		{
			BranchID:     &branchID,
			Amount:       12000,
			UserID:       &userID,
			PurchaseDate: date.AddDate(0, 0, 30),
		},
	}
}

func Test_HasCampaign_Completed(t *testing.T) {
	tx := &modelTransaction.Transaction{
		Amount: 5000,
	}
	startDate, _ := time.Parse("2006-01-02T00:00:00-05:00", "2023-05-15T00:00:00-05:00")
	endDate, _ := time.Parse("2006-01-02T00:00:00-05:00", "2023-05-30T23:59:59-05:00")
	campaign := &modelCampaign.Campaign{
		StartDate: startDate,
		EndDate:   endDate,
		Completed: true,
	}
	apply := util.HasCampaign(tx, campaign)
	assert.False(t, apply)

}

func Test_HasCampaign_MinAmount_False(t *testing.T) {
	tx := &modelTransaction.Transaction{
		Amount: 5000,
	}
	startDate, _ := time.Parse("2006-01-02T00:00:00-05:00", "2023-05-15T00:00:00-05:00")
	endDate, _ := time.Parse("2006-01-02T00:00:00-05:00", "2023-05-30T23:59:59-05:00")
	campaign := &modelCampaign.Campaign{
		StartDate: startDate,
		EndDate:   endDate,
		MinAmount: 20000,
	}
	apply := util.HasCampaign(tx, campaign)
	assert.False(t, apply)

}

func Test_HasCampaign_OutDate(t *testing.T) {
	tx := &modelTransaction.Transaction{
		Amount: 25000,
	}
	startDate, _ := time.Parse("2006-01-02T00:00:00-05:00", "2023-05-15T00:00:00-05:00")
	endDate, _ := time.Parse("2006-01-02T00:00:00-05:00", "2023-05-30T23:59:59-05:00")
	campaign := &modelCampaign.Campaign{
		StartDate: startDate,
		EndDate:   endDate,
		MinAmount: 20000,
	}
	apply := util.HasCampaign(tx, campaign)
	assert.False(t, apply)
}

func Test_HasCampaign_False_Redeemed(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-15T00:00:00")
	endDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-30T23:59:59")
	txDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-20T15:59:59")
	tx := &modelTransaction.Transaction{
		Amount:       25000,
		PurchaseDate: txDate,
		Redeemed:     true,
	}
	campaign := &modelCampaign.Campaign{
		StartDate: startDate,
		EndDate:   endDate,
		MinAmount: 20000,
	}
	apply := util.HasCampaign(tx, campaign)
	assert.False(t, apply)
}

func Test_HasCampaign_False_Accumulated(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-15T00:00:00")
	endDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-30T23:59:59")
	txDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-20T15:59:59")
	tx := &modelTransaction.Transaction{
		Amount:       25000,
		PurchaseDate: txDate,
		Accumulated:  true,
	}
	campaign := &modelCampaign.Campaign{
		StartDate: startDate,
		EndDate:   endDate,
		MinAmount: 20000,
	}
	apply := util.HasCampaign(tx, campaign)
	assert.False(t, apply)
}

func Test_HasCampaign_True(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-15T00:00:00")
	endDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-30T23:59:59")
	txDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-20T15:59:59")
	tx := &modelTransaction.Transaction{
		Amount:       25000,
		PurchaseDate: txDate,
	}
	campaign := &modelCampaign.Campaign{
		StartDate: startDate,
		EndDate:   endDate,
		MinAmount: 20000,
	}
	apply := util.HasCampaign(tx, campaign)
	assert.True(t, apply)
}

func Test_CalcPointsAndCashbackAccumulate_Not_Apply(t *testing.T) {
	transactions := dummyTransactions()
	startDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-15T00:00:00")
	endDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-30T23:59:59")
	campaign := &modelCampaign.Campaign{
		StartDate:        startDate,
		EndDate:          endDate,
		RewardMultiplier: 0.3,
	}
	points, cashback := util.CalcPointsAndCashbackAccumulate(transactions[0], campaign, 0.001, 0.001)
	assert.Equal(t, points, int(0))
	assert.Equal(t, cashback, float64(0))

}

func Test_CalcPointsAndCashbackAccumulate_Apply(t *testing.T) {
	transactions := dummyTransactions()
	startDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-15T00:00:00")
	endDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-30T23:59:59")
	campaign := &modelCampaign.Campaign{
		StartDate:        startDate,
		EndDate:          endDate,
		RewardMultiplier: 1,
	}
	points, cashback := util.CalcPointsAndCashbackAccumulate(transactions[1], campaign, 0.001, 0.001)
	assert.Equal(t, int(23), points)
	assert.Equal(t, float64(23), cashback)

}

func Test_CalcPointsAndCashbackAccumulate_Apply_2(t *testing.T) {
	transactions := dummyTransactions()
	startDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-15T00:00:00")
	endDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-30T23:59:59")
	campaign := &modelCampaign.Campaign{
		StartDate:        startDate,
		EndDate:          endDate,
		RewardMultiplier: 0.3,
		MinAmount:        20000,
	}
	points, cashback := util.CalcPointsAndCashbackAccumulate(transactions[1], campaign, 0.001, 0.001)
	assert.Equal(t, int(6), points)
	assert.Equal(t, "6.9", fmt.Sprintf("%.1f", cashback))

}

func Test_CalcPointsAndCashbackAccumulate_Apply_3(t *testing.T) {
	transactions := dummyTransactions()
	startDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-15T00:00:00")
	endDate, _ := time.Parse("2006-01-02T15:04:05", "2023-05-30T23:59:59")
	campaign := &modelCampaign.Campaign{
		StartDate:        startDate,
		EndDate:          endDate,
		RewardMultiplier: 1,
	}
	points, cashback := util.CalcPointsAndCashbackAccumulate(transactions[1], campaign, 0.001, 0.001)
	assert.Equal(t, int(23), points)
	assert.Equal(t, float64(23), cashback)

}
