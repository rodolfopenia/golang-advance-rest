package repository

import (
	"context"

	"github.com/rodolfopenia/go-advance-rest/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	Close() error
}

// El codigo debe basarse en abstracciones
// y no en cosas concretas

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}
