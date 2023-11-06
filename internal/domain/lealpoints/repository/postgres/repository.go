package postgres

import (
	"github.com/braejan/go-leal-challenge/internal/domain/lealpoints/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresLealPointRepository struct {
	db *gorm.DB
}

func NewPostgresLealPointRepository(db *gorm.DB) *postgresLealPointRepository {
	return &postgresLealPointRepository{db}
}

func (r *postgresLealPointRepository) GetLealPointByBusinessID(ID uuid.UUID) (lealPoint *model.LealPoint, err error) {
	if ID == uuid.Nil {
		return nil, nil
	}
	err = r.db.First(&lealPoint, "business_id = ?", ID).Error
	return
}
