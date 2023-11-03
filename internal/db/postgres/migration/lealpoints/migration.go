package lealpoints

import (
	"context"
	"errors"

	businessModel "github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	lealpointModel "github.com/braejan/go-leal-challenge/internal/domain/lealpoints/model"
	"gorm.io/gorm"
)

func LealPointMigrationData(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&lealpointModel.LealPoint{}) {
		if err = db.First(&lealpointModel.LealPoint{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			ctx := context.Background()
			business := &businessModel.Business{}
			err = db.First(business).Error
			if err != nil {
				return
			}
			lealPoint := &lealpointModel.LealPoint{
				BusinessID:       business.ID,
				ConversionFactor: 0.001,
			}
			err = db.WithContext(ctx).Create(lealPoint).Error
		}
	}
	return
}
