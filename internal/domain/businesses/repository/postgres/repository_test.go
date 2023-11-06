package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	businessPostgresRepo "github.com/braejan/go-leal-challenge/internal/domain/businesses/repository/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CreateBusiness_Err_Nil(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	var business *model.Business
	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	err := repo.CreateBusiness(context.Background(), business)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)

}

func Test_CreateBusiness_Error_Insert(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"businesses\" (.+) ").WillReturnError(gorm.ErrInvalidDB)
	mock.ExpectRollback()
	business := &model.Business{}
	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	err := repo.CreateBusiness(context.Background(), business)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_GetBusinessByID_Invalid_ID(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	_, err := repo.GetBusinessByID(context.Background(), uuid.Nil)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_GetBusinessByID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"})
	expected.AddRow(ID, time.Time{}, time.Time{}, nil, "mi primera chamba")
	mock.ExpectQuery("SELECT .+ FROM \"businesses\" .+").WithArgs(&ID).WillReturnRows(expected)
	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	business, err := repo.GetBusinessByID(context.Background(), ID)
	assert.Nil(t, err)
	assert.Equal(t, ID.String(), business.ID.String())
}

func Test_GetBusinessByID_Error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	mock.ExpectQuery("SELECT .+ FROM \"businesses\" .+").WithArgs(&ID).WillReturnError(gorm.ErrRecordNotFound)
	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	business, err := repo.GetBusinessByID(context.Background(), ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, business)
}

func Test_UpdateBusiness_Error_Nil(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	var business *model.Business
	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	err := repo.UpdateBusiness(context.Background(), business)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_UpdateBusiness_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	business := &model.Business{
		ID:   &ID,
		Name: "mi primera chamba",
	}
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"})
	expected.AddRow(business.ID, business.CreatedAt, business.UpdatedAt, business.DeletedAt, business.Name)
	mock.ExpectQuery("SELECT .+ FROM \"businesses\" .+").WithArgs(&ID).WillReturnRows(expected)

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE \"businesses\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	err := repo.UpdateBusiness(context.Background(), business)
	assert.Nil(t, err)
}

func Test_DeleteBusinessByID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	business := &model.Business{
		ID:   &ID,
		Name: "mi primera chamba",
	}
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"})
	expected.AddRow(business.ID, business.CreatedAt, business.UpdatedAt, business.DeletedAt, business.Name)
	mock.ExpectQuery("SELECT .+ FROM \"businesses\" .+").WithArgs(&ID).WillReturnRows(expected)

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE \"businesses\" SET .+deleted_at.+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	err := repo.DeleteBusinessByID(context.Background(), ID)
	assert.Nil(t, err)
}

func Test_GetBusinessesByBusinessID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"})
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, time.Time{}, "mi primera chamba")
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, time.Time{}, "mi segunda chamba")
	mock.ExpectQuery("SELECT .+ FROM \"businesses\" .+").WillReturnRows(expected)
	repo := businessPostgresRepo.NewPostgresBusinessRepository(db)
	businesses, err := repo.ListBusinesses(context.Background())
	assert.Nil(t, err)
	assert.Len(t, businesses, 2)
}
