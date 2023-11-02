package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	branchModel "github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	businessModel "github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	campaignModel "github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormConfiguration = &gorm.Config{
	//TODO: Not implemented yet
}

type postgresDatasource struct {
	user     string
	password string
	database string
	host     string
	port     string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getDataSourceFromEnv() (datasource *postgresDatasource) {
	datasource = &postgresDatasource{}
	datasource.user = getEnv("POSTGRES_USER", "postgres")
	datasource.password = getEnv("POSTGRES_PASSWORD", "postgres")
	datasource.database = getEnv("POSTGRES_DATABASE", "leal-challenge-database")
	datasource.host = getEnv("POSTGRES_HOST", "localhost")
	datasource.port = getEnv("POSTGRES_PORT", "5432")
	return
}

func OpenDefautDatabase() (db *gorm.DB, err error) {
	datasourceInfo := getDataSourceFromEnv()
	postgresConfig := postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=America/Bogota",
			datasourceInfo.user,
			datasourceInfo.password,
			datasourceInfo.database,
			datasourceInfo.host,
			datasourceInfo.port,
		),
		PreferSimpleProtocol: true,
	}
	db, err = gorm.Open(postgres.New(postgresConfig), gormConfiguration)
	if err != nil {
		return
	}
	log.Println("Database connected")
	log.Println("Running migrations")
	err = migrations(db)
	return
}

func migrations(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		&businessModel.Business{},
		&branchModel.Branch{},
		&campaignModel.Campaign{},
	)
	if err != nil {
		log.Fatal("Finish automigrate with error", err)
	}
	err = businessMigrationData(db)
	if err != nil {
		log.Fatal("business migration error", err)
	}
	err = branchMigrationData(db)
	if err != nil {
		log.Fatal("branch migration error", err)
	}
	return
}

func businessMigrationData(db *gorm.DB) (err error) {
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

func branchMigrationData(db *gorm.DB) (err error) {
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
