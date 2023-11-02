package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       string    `json:"name"`
	BusinessID string    `json:"business_id"`
}
