package repository

import (
	"context"
	"golang-api-standard-http-lib/internal/domain"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
}
