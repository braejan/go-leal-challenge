package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	campaignPostgresRepo "github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CreateCampaign_Err_Nil(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	var campaign *model.Campaign
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	err := repo.CreateCampaign(context.Background(), campaign)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_CreateCampaign_Error_Insert(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"campaigns\" (.+) ").WillReturnError(gorm.ErrInvalidDB)
	mock.ExpectRollback()
	campaign := &model.Campaign{}
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	err := repo.CreateCampaign(context.Background(), campaign)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_GetCampaignByID_Invalid_ID(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	_, err := repo.GetCampaignByID(context.Background(), uuid.Nil)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_GetCampaignByID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "business_id", "branch_id", "name", "start_date", "end_date", "reward_multiplier", "min_amount", "completed"})
	expected.AddRow(ID, time.Time{}, time.Time{}, nil, uuid.New(), uuid.New(), "mi primera chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WithArgs(&ID).WillReturnRows(expected)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	business, err := repo.GetCampaignByID(context.Background(), ID)
	assert.Nil(t, err)
	assert.Equal(t, ID.String(), business.ID.String())
}

func Test_GetCampaignByID_Error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WithArgs(&ID).WillReturnError(gorm.ErrRecordNotFound)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	campaign, err := repo.GetCampaignByID(context.Background(), ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, campaign)
}

func Test_UpdateCampaign_Error_Nil(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	var campaign *model.Campaign
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	err := repo.UpdateCampaign(context.Background(), campaign)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_UpdateCampaign_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	campaign := &model.Campaign{
		ID:   &ID,
		Name: "mi primera chamba",
	}
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "business_id", "branch_id", "name", "start_date", "end_date", "reward_multiplier", "min_amount", "completed"})
	expected.AddRow(ID, time.Time{}, time.Time{}, nil, uuid.New(), uuid.New(), "mi primera chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WithArgs(&ID).WillReturnRows(expected)

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE \"campaigns\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	err := repo.UpdateCampaign(context.Background(), campaign)
	assert.Nil(t, err)
}

func Test_DeleteCampaignByID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "business_id", "branch_id", "name", "start_date", "end_date", "reward_multiplier", "min_amount", "completed"})
	expected.AddRow(ID, time.Time{}, time.Time{}, nil, uuid.New(), uuid.New(), "mi primera chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WithArgs(&ID).WillReturnRows(expected)

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE \"campaigns\" SET .+deleted_at.+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	err := repo.DeleteCampaignByID(context.Background(), ID)
	assert.Nil(t, err)
}

func Test_ListCampaignsByBusinessID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "business_id", "branch_id", "name", "start_date", "end_date", "reward_multiplier", "min_amount", "completed"})
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, nil, ID, uuid.New(), "mi primera chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, nil, ID, uuid.New(), "mi segunda chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WillReturnRows(expected)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	campaigns, err := repo.ListCampaignsByBusinessID(context.Background(), ID)
	assert.Nil(t, err)
	assert.Len(t, campaigns, 2)
	assert.Equal(t, ID.String(), campaigns[0].BusinessID.String())
}

func Test_ListCampaignsByBusinessID_Err_ID_Nil(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.Nil
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WillReturnError(gorm.ErrInvalidValue)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	_, err := repo.ListCampaignsByBusinessID(context.Background(), ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_ListCampaignsByBranchID_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.New()
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "business_id", "branch_id", "name", "start_date", "end_date", "reward_multiplier", "min_amount", "completed"})
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, nil, ID, uuid.New(), "mi primera chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, nil, ID, uuid.New(), "mi segunda chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WillReturnRows(expected)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	campaigns, err := repo.ListCampaignsByBranchID(context.Background(), ID)
	assert.Nil(t, err)
	assert.Len(t, campaigns, 2)
	assert.Equal(t, ID.String(), campaigns[0].BusinessID.String())
}

func Test_ListCampaignsByBranchID_Err_ID_Nil(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ID := uuid.Nil
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WillReturnError(gorm.ErrInvalidValue)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	_, err := repo.ListCampaignsByBranchID(context.Background(), ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_GetUncompletedCampaigns_Err(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WillReturnError(gorm.ErrInvalidValue)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	_, err := repo.GetUncompletedCampaigns(context.Background())
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_GetUncompletedCampaigns_Sucess(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	expected := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "business_id", "branch_id", "name", "start_date", "end_date", "reward_multiplier", "min_amount", "completed"})
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, nil, uuid.New(), uuid.New(), "mi primera chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	expected.AddRow(uuid.New(), time.Time{}, time.Time{}, nil, uuid.New(), uuid.New(), "mi segunda chamba", time.Time{}, time.Time{}, 1.3, 20000, false)
	mock.ExpectQuery("SELECT .+ FROM \"campaigns\" .+").WillReturnRows(expected)
	repo := campaignPostgresRepo.NewPostgresCampaignRepository(db)
	campaigns, err := repo.GetUncompletedCampaigns(context.Background())
	assert.Nil(t, err)
	assert.Len(t, campaigns, 2)
}
