package model

import (
	transactionModel "github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model   `json:"-"`
	ID           *uuid.UUID                     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string                         `json:"name"`
	LealPoints   int64                          `json:"leal_points"`
	LealCoins    float64                        `json:"leal_coins"`
	Transactions []transactionModel.Transaction `json:"-"`
}
