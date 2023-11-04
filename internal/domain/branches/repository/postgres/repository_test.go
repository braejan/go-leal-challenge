package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	branchPostgresRepo "github.com/braejan/go-leal-challenge/internal/domain/branches/repository/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CreateBranch_Error_Nil(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	var branch *model.Branch
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	err := repo.CreateBranch(context.Background(), branch)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_CreateBranch_Error_Insert(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"branches\" (.+) ").WillReturnError(gorm.ErrInvalidDB)
	mock.ExpectRollback()
	branch := &model.Branch{}
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	err := repo.CreateBranch(context.Background(), branch)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_GetBranchByID_Invalid_ID(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	_, err := repo.GetBranchByID(context.Background(), uuid.Nil)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_GetBranchByID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "business_id"})
	expected.AddRow(ID, time.Time{}, time.Time{}, time.Time{}, "mi primera chamba", uuid.New())
	mock.ExpectQuery("SELECT .+ FROM \"branches\" .+").WithArgs(&ID).WillReturnRows(expected)
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	branch, err := repo.GetBranchByID(context.Background(), ID)
	assert.Nil(t, err)
	assert.Equal(t, ID.String(), branch.ID.String())
}

func Test_GetBranchByID_Error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	mock.ExpectQuery("SELECT .+ FROM \"branches\" .+").WithArgs(&ID).WillReturnError(gorm.ErrRecordNotFound)
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	branch, err := repo.GetBranchByID(context.Background(), ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, branch)
}

func Test_GetBranchesByBusinessID_Invalid_ID(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	_, err := repo.GetBranchesByBusinessID(context.Background(), uuid.Nil)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_GetBranchesByBusinessID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "business_id"})
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, time.Time{}, "mi primera chamba", ID)
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, time.Time{}, "mi segunda chamba", ID)
	mock.ExpectQuery("SELECT .+ FROM \"branches\" .+").WithArgs(&ID).WillReturnRows(expected)
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	branches, err := repo.GetBranchesByBusinessID(context.Background(), ID)
	assert.Nil(t, err)
	assert.Len(t, branches, 2)
}

func Test_GetBranchesByBusinessID_Error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	mock.ExpectQuery("SELECT .+ FROM \"branches\" .+").WithArgs(&ID).WillReturnError(gorm.ErrRecordNotFound)
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	branches, err := repo.GetBranchesByBusinessID(context.Background(), ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, branches)
}

func Test_UpdateBranch_Error_Nil(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	var branch *model.Branch
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	err := repo.UpdateBranch(context.Background(), branch)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_UpdateBranch_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	branch := &model.Branch{
		ID:         &ID,
		Name:       "mi primera chamba",
		BusinessID: &ID,
	}
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "business_id"})
	expected.AddRow(branch.ID, branch.CreatedAt, branch.UpdatedAt, branch.DeletedAt, branch.Name, branch.BusinessID)
	mock.ExpectQuery("SELECT .+ FROM \"branches\" .+").WithArgs(&ID).WillReturnRows(expected)

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE \"branches\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	err := repo.UpdateBranch(context.Background(), branch)
	assert.Nil(t, err)
}

func Test_DeleteBranch_Error_Nil(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	err := repo.DeleteBranchByID(context.Background(), uuid.Nil)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_DeleteBranchByID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	branch := &model.Branch{
		ID:         &ID,
		Name:       "mi primera chamba",
		BusinessID: &ID,
	}
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "business_id"})
	expected.AddRow(branch.ID, branch.CreatedAt, branch.UpdatedAt, branch.DeletedAt, branch.Name, branch.BusinessID)
	mock.ExpectQuery("SELECT .+ FROM \"branches\" .+").WithArgs(&ID).WillReturnRows(expected)

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE \"branches\" SET .+deleted_at.+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := branchPostgresRepo.NewPostgresBranchRepository(db)
	err := repo.DeleteBranchByID(context.Background(), ID)
	assert.Nil(t, err)
}
