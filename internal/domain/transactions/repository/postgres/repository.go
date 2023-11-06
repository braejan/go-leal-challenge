package postgres

import (
	"github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
	"github.com/braejan/go-leal-challenge/internal/domain/transactions/repository"
	"gorm.io/gorm"
)

type postrgresTransactionRepository struct {
	db *gorm.DB
}

func NewPostgresTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &postrgresTransactionRepository{db}
}

func (r *postrgresTransactionRepository) GetTransactionsToProcess() (transactions []*model.Transaction, err error) {
	err = r.db.Where("accumulated = ? AND redeemed = ?", false, false).Find(&transactions).Error
	return
}
