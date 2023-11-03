package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model       `json:"-"`
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BusinessID       uuid.UUID `json:"business_id" gorm:"type:uuid;not null"`
	BranchID         uuid.UUID `json:"branch_id" gorm:"type:uuid"`
	Name             string    `json:"name" gorm:"not null"`
	StartDate        time.Time `json:"start_date" gorm:"not null"`
	EndDate          time.Time `json:"end_date" gorm:"not null"`
	RewardMultiplier float64   `json:"reward_multiplier" gorm:"not null"`
}
