package repository

import "github.com/braejan/go-leal-challenge/internal/domain/transactions/model"

type TransactionRepository interface {
	GetTransactionsToProcess() ([]*model.Transaction, error)
}
