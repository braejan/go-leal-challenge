package users

import (
	"context"
	"errors"

	userModel "github.com/braejan/go-leal-challenge/internal/domain/users/model"
	"gorm.io/gorm"
)

func UserMigrationData(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&userModel.User{}) {
		if err = db.First(&userModel.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			ctx := context.Background()
			user := &userModel.User{
				Name:       "José Miguel",
				LealPoints: 0,
				LealCoins:  0.0,
			}
			err = db.WithContext(ctx).Create(user).Error
		}
	}
	return
}
