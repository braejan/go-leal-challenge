package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
}
