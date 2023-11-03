package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	CreateLealPoint(ctx context.Context, lealPoint *model.Transaction) error
	GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *model.Transaction) error
	DeleteTransaction(ctx context.Context, id uuid.UUID) error
}
