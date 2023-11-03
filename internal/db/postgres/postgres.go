package postgres

import (
	"fmt"
	"log"

	"github.com/braejan/go-leal-challenge/internal/db/postgres/migration/branches"
	"github.com/braejan/go-leal-challenge/internal/db/postgres/migration/businesses"
	"github.com/braejan/go-leal-challenge/internal/db/postgres/migration/lealcoins"
	"github.com/braejan/go-leal-challenge/internal/db/postgres/migration/lealpoints"
	"github.com/braejan/go-leal-challenge/internal/db/postgres/migration/transactions"
	"github.com/braejan/go-leal-challenge/internal/db/postgres/migration/users"
	branchModel "github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	businessModel "github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	campaignModel "github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	lealcoinsModel "github.com/braejan/go-leal-challenge/internal/domain/lealcoins/model"
	lealpointsModel "github.com/braejan/go-leal-challenge/internal/domain/lealpoints/model"
	transactionsModel "github.com/braejan/go-leal-challenge/internal/domain/transactions/model"
	userModel "github.com/braejan/go-leal-challenge/internal/domain/users/model"
	"github.com/braejan/go-leal-challenge/pkg/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormConfiguration = &gorm.Config{}

type postgresDatasource struct {
	user     string
	password string
	database string
	host     string
	port     string
}

func getDataSourceFromEnv() (datasource *postgresDatasource) {
	datasource = &postgresDatasource{}
	datasource.user = util.GetEnv("POSTGRES_USER", "postgres")
	datasource.password = util.GetEnv("POSTGRES_PASSWORD", "postgres")
	datasource.database = util.GetEnv("POSTGRES_DATABASE", "leal-challenge-database")
	datasource.host = util.GetEnv("POSTGRES_HOST", "localhost")
	datasource.port = util.GetEnv("POSTGRES_PORT", "5432")
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
		&userModel.User{},
		&transactionsModel.Transaction{},
		&lealcoinsModel.LealCoin{},
		&lealpointsModel.LealPoint{},
	)
	if err != nil {
		log.Fatal("Finish automigrate with error: ", err)
	}
	err = businesses.BusinessMigrationData(db)
	if err != nil {
		log.Fatal("business migration error: ", err)
	}
	err = branches.BranchMigrationData(db)
	if err != nil {
		log.Fatal("branch migration error: ", err)
	}
	err = users.UserMigrationData(db)
	if err != nil {
		log.Fatal("user migration error: ", err)
	}
	err = transactions.TransactionsMigrationData(db)
	if err != nil {
		log.Fatal("transactions migration error: ", err)
	}
	err = lealcoins.LealCoinMigrationData(db)
	if err != nil {
		log.Fatal("lealcoins migration error: ", err)
	}
	err = lealpoints.LealPointMigrationData(db)
	if err != nil {
		log.Fatal("lealpoints migration error: ", err)
	}
	return
}
