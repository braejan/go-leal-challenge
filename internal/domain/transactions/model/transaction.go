package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model   `json:"-"`
	ID           *uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BranchID     *uuid.UUID `json:"branch_id" gorm:"type:uuid;not null"`
	UserID       *uuid.UUID `json:"user_id" gorm:"not null"`
	Amount       float64    `json:"amount" gorm:"not null"`
	PurchaseDate time.Time  `json:"purchase_date" gorm:"not null"`
	Accumulated  bool       `json:"accumulated" gorm:"not null;default:false;index"`
	Redeemed     bool       `json:"redeemed" gorm:"not null;default:false;index"`
}
