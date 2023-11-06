package postgres

import (
	"github.com/braejan/go-leal-challenge/internal/domain/lealcoins/model"
	"github.com/braejan/go-leal-challenge/internal/domain/lealcoins/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresLealCoinRepository struct {
	db *gorm.DB
}

func NewPostgresLealCoinRepository(db *gorm.DB) repository.LealCoinRepository {
	return &postgresLealCoinRepository{db: db}
}

func (r *postgresLealCoinRepository) GetLealCoinByBusinessID(ID uuid.UUID) (lealCoin *model.LealCoin, err error) {
	if ID == uuid.Nil {
		err = gorm.ErrInvalidValue
		return
	}
	err = r.db.First(&lealCoin, "business_id = ?", ID).Error
	return
}
