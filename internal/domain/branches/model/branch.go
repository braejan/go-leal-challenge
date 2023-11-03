package model

import (
	campaignModel "github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	transactionModel "github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model   `json:"-"`
	ID           uuid.UUID                      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string                         `json:"name" gorm:"not null"`
	BusinessID   uuid.UUID                      `json:"business_id" gorm:"type:uuid;not null"`
	Campaigns    []campaignModel.Campaign       `json:"-"`
	Transactions []transactionModel.Transaction `json:"-"`
}
