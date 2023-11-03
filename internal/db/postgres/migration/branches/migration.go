package branches

import (
	"context"
	"errors"

	branchModel "github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	businessModel "github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"gorm.io/gorm"
)

func BranchMigrationData(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&branchModel.Branch{}) {
		if err = db.First(&branchModel.Branch{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			ctx := context.Background()
			business := &businessModel.Business{}
			err = db.First(business).Error
			if err != nil {
				return
			}
			branches := []*branchModel.Branch{
				{
					BusinessID: business.ID,
					Name:       "Sucursal 1",
				},
				{
					BusinessID: business.ID,
					Name:       "Sucursal 2",
				},
			}
			err = db.WithContext(ctx).Create(branches).Error
		}
	}
	return
}
