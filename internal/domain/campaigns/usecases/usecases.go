package usecases

import (
	"context"
	"log"

	branchRepository "github.com/braejan/go-leal-challenge/internal/domain/branches/repository"
	branchPostgres "github.com/braejan/go-leal-challenge/internal/domain/branches/repository/postgres"
	businessRepository "github.com/braejan/go-leal-challenge/internal/domain/businesses/repository"
	businessPostgres "github.com/braejan/go-leal-challenge/internal/domain/businesses/repository/postgres"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	campaignRepository "github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository"
	campaignPostgres "github.com/braejan/go-leal-challenge/internal/domain/campaigns/repository/postgres"
	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/util"
	lealcoinRepository "github.com/braejan/go-leal-challenge/internal/domain/lealcoins/repository"
	lealcoinPostgres "github.com/braejan/go-leal-challenge/internal/domain/lealcoins/repository/postgres"
	lealPointRepository "github.com/braejan/go-leal-challenge/internal/domain/lealpoints/repository"
	lealPointPostgres "github.com/braejan/go-leal-challenge/internal/domain/lealpoints/repository/postgres"
	transactionRepository "github.com/braejan/go-leal-challenge/internal/domain/transactions/repository"
	transactionPostgres "github.com/braejan/go-leal-challenge/internal/domain/transactions/repository/postgres"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignUsecases interface {
	GetCampaignByID(ID uuid.UUID) (model.Campaign, error)
	CreateNewCampaign(campaign model.Campaign) error
	GetCampaignsByBusinessID(ID uuid.UUID) ([]model.Campaign, error)
	GetCampaignsByBranchID(ID uuid.UUID) ([]model.Campaign, error)
	AccumulatePointsAndCashBack() (resume map[uuid.UUID]*model.Accumulated, err error)
}

type campaingUsecases struct {
	campaignRepository.CampaignRepository
	businessRepository.BusinessRepository
	branchRepository.BranchRepository
	transactionRepository.TransactionRepository
	lealcoinRepository.LealCoinRepository
	lealPointRepository.LealPointRepository
}

func NewCampaignUsecases(db *gorm.DB) CampaignUsecases {
	repoCampaign := campaignPostgres.NewPostgresCampaignRepository(db)
	repoBusiness := businessPostgres.NewPostgresBusinessRepository(db)
	repoBranch := branchPostgres.NewPostgresBranchRepository(db)
	repoTransaction := transactionPostgres.NewPostgresTransactionRepository(db)
	repoLealCoin := lealcoinPostgres.NewPostgresLealCoinRepository(db)
	repoLealPoint := lealPointPostgres.NewPostgresLealPointRepository(db)
	return NewCampaignUsecasesWithRepo(repoCampaign, repoBusiness, repoBranch, repoTransaction, repoLealCoin, repoLealPoint)
}
func NewCampaignUsecasesWithRepo(
	repoCampaign campaignRepository.CampaignRepository,
	repoBusiness businessRepository.BusinessRepository,
	repoBranch branchRepository.BranchRepository,
	repoTransaction transactionRepository.TransactionRepository,
	repoLealCoin lealcoinRepository.LealCoinRepository,
	repoLealPoint lealPointRepository.LealPointRepository,
) CampaignUsecases {
	return &campaingUsecases{
		CampaignRepository:    repoCampaign,
		BusinessRepository:    repoBusiness,
		BranchRepository:      repoBranch,
		TransactionRepository: repoTransaction,
		LealCoinRepository:    repoLealCoin,
		LealPointRepository:   repoLealPoint,
	}
}

func (u *campaingUsecases) GetCampaignByID(ID uuid.UUID) (campaign model.Campaign, err error) {
	found, err := u.CampaignRepository.GetCampaignByID(context.Background(), ID)
	if err != nil {
		return
	}
	campaign = *found
	return
}
func (u *campaingUsecases) CreateNewCampaign(campaign model.Campaign) (err error) {
	found, err := u.BranchRepository.GetBranchByID(context.Background(), *campaign.BranchID)
	if err != nil {
		return
	}
	if *found.BusinessID != *campaign.BusinessID {
		err = gorm.ErrInvalidValue
		return
	}
	campaign.Completed = false
	return u.CampaignRepository.CreateCampaign(context.Background(), &campaign)
}
func (u *campaingUsecases) GetCampaignsByBusinessID(ID uuid.UUID) (campaigns []model.Campaign, err error) {
	_, err = u.BusinessRepository.GetBusinessByID(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	found, err := u.CampaignRepository.ListCampaignsByBusinessID(context.Background(), ID)
	if err != nil {
		return
	}
	campaigns = make([]model.Campaign, len(found))
	for i, c := range found {
		campaigns[i] = *c
	}
	return
}
func (u *campaingUsecases) GetCampaignsByBranchID(ID uuid.UUID) (campaigns []model.Campaign, err error) {
	_, err = u.BranchRepository.GetBranchByID(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	found, err := u.CampaignRepository.ListCampaignsByBranchID(context.Background(), ID)
	if err != nil {
		return
	}
	campaigns = make([]model.Campaign, len(found))
	for i, c := range found {
		campaigns[i] = *c
	}
	return
}

func (u *campaingUsecases) AccumulatePointsAndCashBack() (resume map[uuid.UUID]*model.Accumulated, err error) {
	transactions, err := u.TransactionRepository.GetTransactionsToProcess()
	if err != nil {
		return
	}
	resume = make(map[uuid.UUID]*model.Accumulated)
	for _, transaction := range transactions {
		// Get the Branch information
		branch, errProcess := u.BranchRepository.GetBranchByID(context.Background(), *transaction.BranchID)
		if errProcess != nil {
			err = errProcess
			return
		}
		// Get the ConversionFactor for points
		lealCoin, errProcess := u.LealCoinRepository.GetLealCoinByBusinessID(*branch.BusinessID)
		if errProcess != nil {
			err = errProcess
			return
		}
		// Get the Equivalence for cashback
		lealPoint, errProcess := u.LealPointRepository.GetLealPointByBusinessID(*branch.BusinessID)
		if errProcess != nil {
			err = errProcess
			return
		}
		points := int(transaction.Amount * lealPoint.ConversionFactor)
		cashBack := transaction.Amount * lealCoin.Equivalence
		if value, ok := resume[*transaction.UserID]; ok {
			value.TotalPoints = value.TotalPoints + points
			value.TotalCashback = value.TotalCashback + cashBack
			resume[*transaction.UserID] = value
		} else {
			acc := &model.Accumulated{
				TotalPoints:   points,
				TotalCashback: cashBack,
			}
			resume[*transaction.UserID] = acc
		}
		// validate if apply any campaign
		campaigns, errProcess := u.GetUnfinishedCampaigns(context.Background(), transaction.PurchaseDate, *branch.ID)
		if errProcess != nil {
			err = errProcess
			return
		}
		for _, campaign := range campaigns {
			pointsCampaign, cashbackCampaign := util.CalcPointsAndCashbackAccumulate(transaction, campaign, lealPoint.ConversionFactor, lealCoin.Equivalence)
			if pointsCampaign > 0 || cashbackCampaign > 0 {
				log.Println("user_id", transaction.UserID.String(), "point extra", pointsCampaign, "cashback extra", cashbackCampaign, "campaign start date", campaign.StartDate)
				resume[*transaction.UserID].TotalPoints = resume[*transaction.UserID].TotalPoints + pointsCampaign
				resume[*transaction.UserID].TotalCashback = resume[*transaction.UserID].TotalCashback + cashbackCampaign
			}
		}
	}
	return
}
