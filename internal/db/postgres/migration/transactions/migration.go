package transactions

import (
	"errors"
	"time"

	branchModel "github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	transactionsModel "github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
	usersModel "github.com/braejan/go-leal-challenge/internal/domain/users/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func generateTransactions(branchID *uuid.UUID, userID *uuid.UUID) []*transactionsModel.Transaction {
	date, _ := time.Parse("2006-01-02 15:04", "2023-05-01 00:00")
	return []*transactionsModel.Transaction{
		{
			BranchID:     branchID,
			Amount:       5000,
			UserID:       userID,
			PurchaseDate: date.AddDate(0, 0, 1),
		},
		{
			BranchID:     branchID,
			Amount:       23000,
			UserID:       userID,
			PurchaseDate: date.AddDate(0, 0, 15),
		},
		{
			BranchID:     branchID,
			Amount:       25000,
			UserID:       userID,
			PurchaseDate: date.AddDate(0, 0, 25),
		},
		{
			BranchID:     branchID,
			Amount:       12000,
			UserID:       userID,
			PurchaseDate: date.AddDate(0, 0, 30),
		},
	}
}

func TransactionsMigrationData(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&transactionsModel.Transaction{}) {
		if err = db.First(&transactionsModel.Transaction{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

			user := &usersModel.User{}
			err = db.First(user).Error
			if err != nil {
				return
			}
			var branches []*branchModel.Branch
			err = db.Find(&branches).Error
			if err != nil {
				return
			}
			for _, branch := range branches {
				err = db.Create(generateTransactions(branch.ID, user.ID)).Error
				if err != nil {
					return
				}
			}
		}
	}
	return
}
