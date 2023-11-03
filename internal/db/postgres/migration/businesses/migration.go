package businesses

import (
	"context"
	"errors"

	businessModel "github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"gorm.io/gorm"
)

func BusinessMigrationData(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&businessModel.Business{}) {
		if err = db.First(&businessModel.Business{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			ctx := context.Background()
			business := &businessModel.Business{
				Name: "Texaco",
			}
			err = db.WithContext(ctx).Create(business).Error
		}
	}
	return
}
