package repository

import (
	"context"

	"github.com/braejan/go-leal-challenge/internal/domain/users/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ListUsers(ctx context.Context) ([]*model.User, error)
}
