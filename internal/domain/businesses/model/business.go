package model

import (
	branchesModel "github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	campaignModel "github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	lealcoinModel "github.com/braejan/go-leal-challenge/internal/domain/lealcoins/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model `json:"-"`
	ID         *uuid.UUID                `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       string                    `json:"name" gorm:"not null"`
	Branches   []*branchesModel.Branch   `json:"-"`
	Campaigns  []*campaignModel.Campaign `json:"-"`
	LealCoin   *lealcoinModel.LealCoin   `json:"-"`
}
