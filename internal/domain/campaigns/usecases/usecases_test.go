package usecases_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	branchModel "github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	mockBranch "github.com/braejan/go-leal-challenge/internal/domain/branches/repository/postgres/mock"
	mockBusiness "github.com/braejan/go-leal-challenge/internal/domain/businesses/repository/postgres/mock"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	mockCampaign "github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository/postgres/mock"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_NewCampaignUsecases_Sucess(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	uc := usecases.NewCampaignUsecases(db)
	assert.NotNil(t, uc)
}
func Test_CreateNewCampaign_Err_BranchRepository_GetBranchByID(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	repoBranch.On("GetBranchByID", testifymock.Anything, testifymock.Anything).Return(nil, gorm.ErrInvalidDB)
	ID := uuid.New()
	campaign := model.Campaign{
		BranchID: &ID,
	}
	err := usecases.CreateNewCampaign(campaign)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidDB, err)
}

func Test_CreateNewCampaign_Err_Diff_Business_ID(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	otherBusinessID := uuid.New()
	repoBranch.On("GetBranchByID", testifymock.Anything, testifymock.Anything).Return(&branchModel.Branch{BusinessID: &otherBusinessID}, nil)
	ID := uuid.New()
	businessID := uuid.New()
	campaign := model.Campaign{
		BranchID:   &ID,
		BusinessID: &businessID,
	}
	err := usecases.CreateNewCampaign(campaign)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_CreateNewCampaign_Err_Create(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	businessID := uuid.New()
	repoBranch.On("GetBranchByID", testifymock.Anything, testifymock.Anything).Return(&branchModel.Branch{BusinessID: &businessID}, nil)
	ID := uuid.New()

	campaign := model.Campaign{
		BranchID:   &ID,
		BusinessID: &businessID,
	}
	repoCampaign.On("CreateCampaign", testifymock.Anything, testifymock.Anything).Return(gorm.ErrInvalidValue)
	err := usecases.CreateNewCampaign(campaign)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrInvalidValue, err)
}

func Test_CreateNewCampaign_Sucess(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	businessID := uuid.New()
	repoBranch.On("GetBranchByID", testifymock.Anything, testifymock.Anything).Return(&branchModel.Branch{BusinessID: &businessID}, nil)
	ID := uuid.New()

	campaign := model.Campaign{
		BranchID:   &ID,
		BusinessID: &businessID,
	}
	repoCampaign.On("CreateCampaign", testifymock.Anything, testifymock.Anything).Return(nil)
	err := usecases.CreateNewCampaign(campaign)
	assert.Nil(t, err)
}

func Test_GetCampaignByID_Err(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	repoCampaign.On("GetCampaignByID", testifymock.Anything, testifymock.Anything).Return(nil, gorm.ErrRecordNotFound)
	_, err := usecases.GetCampaignByID(uuid.New())
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_GetCampaignByID_Sucess(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	campaign := model.Campaign{
		Name: "mi primera chamba",
	}
	repoCampaign.On("GetCampaignByID", testifymock.Anything, testifymock.Anything).Return(&campaign, nil)
	found, err := usecases.GetCampaignByID(uuid.New())
	assert.Nil(t, err)
	assert.Equal(t, "mi primera chamba", found.Name)
}

func Test_GetCampaignsByBusinessID_Err_GetBusinessByID(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	ID := uuid.New()
	repoBusiness.On("GetBusinessByID", testifymock.Anything, ID).Return(nil, gorm.ErrRecordNotFound)
	_, err := usecases.GetCampaignsByBusinessID(ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_GetCampaignsByBusinessID_Err_ListCampaignsByBusinessID(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	ID := uuid.New()
	repoBusiness.On("GetBusinessByID", testifymock.Anything, ID).Return(nil, nil)
	repoCampaign.On("ListCampaignsByBusinessID", testifymock.Anything, ID).Return(nil, gorm.ErrRecordNotFound)
	_, err := usecases.GetCampaignsByBusinessID(ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_GetCampaignsByBusinessID_Sucess(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	ID := uuid.New()
	repoBusiness.On("GetBusinessByID", testifymock.Anything, ID).Return(nil, nil)
	campaigns := []*model.Campaign{
		{
			ID:   &ID,
			Name: "mi primera chamba",
		},
		{
			ID:   &ID,
			Name: "mi segunda chamba",
		},
	}
	repoCampaign.On("ListCampaignsByBusinessID", testifymock.Anything, ID).Return(campaigns, nil)
	found, err := usecases.GetCampaignsByBusinessID(ID)
	assert.Nil(t, err)
	assert.Len(t, found, 2)
	assert.Equal(t, "mi primera chamba", found[0].Name)
	assert.Equal(t, "mi segunda chamba", found[1].Name)
}

func Test_GetCampaignsByBranchID_Err_GetBranchByID(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	ID := uuid.New()
	repoBranch.On("GetBranchByID", testifymock.Anything, ID).Return(nil, gorm.ErrRecordNotFound)
	_, err := usecases.GetCampaignsByBranchID(ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_GetCampaignsByBranchID_Err_ListCampaignsByBranchID(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	ID := uuid.New()
	repoBranch.On("GetBranchByID", testifymock.Anything, ID).Return(nil, nil)
	repoCampaign.On("ListCampaignsByBranchID", testifymock.Anything, ID).Return(nil, gorm.ErrRecordNotFound)
	_, err := usecases.GetCampaignsByBranchID(ID)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func Test_GetCampaignsByBranchID_Sucess(t *testing.T) {
	repoCampaign := mockCampaign.NewMockCampaignRepository()
	repoBusiness := mockBusiness.NewMockBusinessRepository()
	repoBranch := mockBranch.NewMockBranchRepository()
	usecases := usecases.NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch)
	ID := uuid.New()
	repoBranch.On("GetBranchByID", testifymock.Anything, ID).Return(nil, nil)
	campaigns := []*model.Campaign{
		{
			ID:   &ID,
			Name: "mi primera chamba",
		},
		{
			ID:   &ID,
			Name: "mi segunda chamba",
		},
	}
	repoCampaign.On("ListCampaignsByBranchID", testifymock.Anything, ID).Return(campaigns, nil)
	found, err := usecases.GetCampaignsByBranchID(ID)
	assert.Nil(t, err)
	assert.Len(t, found, 2)
	assert.Equal(t, "mi primera chamba", found[0].Name)
	assert.Equal(t, "mi segunda chamba", found[1].Name)
}
