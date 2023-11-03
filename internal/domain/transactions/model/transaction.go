package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model    `json:"-"`
	ID            *uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BranchID      *uuid.UUID `json:"branch_id" gorm:"type:uuid;not null"`
	UserID        *uuid.UUID `json:"user_id" gorm:"not null"`
	Amount        int        `json:"amount" gorm:"not null"`
	PurchaseDate  time.Time  `json:"purchase_date" gorm:"not null"`
	RedeemedPoint bool       `json:"redeemed_point" gorm:"not null;default:false;index"`
	RedeemedCoins bool       `json:"redeemed_coins" gorm:"not null;default:false;index"`
}
