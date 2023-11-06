package repository

import (
	"github.com/braejan/go-leal-challenge/internal/domain/lealpoints/model"
	"github.com/google/uuid"
)

type LealPointRepository interface {
	GetLealPointByBusinessID(ID uuid.UUID) (*model.LealPoint, error)
}
