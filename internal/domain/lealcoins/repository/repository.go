package repository

import (
	"github.com/braejan/go-leal-challenge/internal/domain/lealcoins/model"
	"github.com/google/uuid"
)

type LealCoinRepository interface {
	GetLealCoinByBusinessID(ID uuid.UUID) (*model.LealCoin, error)
}
