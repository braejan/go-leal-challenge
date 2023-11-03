package lealcoins

import (
	"context"
	"errors"

	businessModel "github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	lealcoinsModel "github.com/braejan/go-leal-challenge/internal/domain/lealcoins/model"
	"gorm.io/gorm"
)

func LealCoinMigrationData(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&lealcoinsModel.LealCoin{}) {
		if err = db.First(&lealcoinsModel.LealCoin{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			ctx := context.Background()
			business := &businessModel.Business{}
			err = db.First(business).Error
			if err != nil {
				return
			}
			lealcoin := &lealcoinsModel.LealCoin{
				BusinessID:  business.ID,
				Equivalence: 0.001,
			}
			err = db.WithContext(ctx).Create(lealcoin).Error
		}
	}
	return
}
