package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LealPoint struct {
	gorm.Model       `json:"-"`
	ID               *uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BusinessID       *uuid.UUID `json:"business_id" gorm:"type:uuid;not null"`
	ConversionFactor float64    `json:"conversion_factor" gorm:"not null"`
}
