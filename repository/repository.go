package repository

import (
	"context"

	"github.com/rodolfopenia/go-advance-rest/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post, userId string) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPost(ctx context.Context, page uint64) ([]*models.Post, error)
	Close() error
}

// El codigo debe basarse en abstracciones
// y no en cosas concretas

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post, userId string) error {
	return implementation.UpdatePost(ctx, post, userId)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func ListPost(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implementation.ListPost(ctx, page)
}
