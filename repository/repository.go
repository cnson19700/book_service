package repository

import (
	"context"

	"github.com/cnson19700/book_service/repository/book"
	"gorm.io/gorm"
)

type Repository struct {
	Book book.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		Book: book.NewPGRepository(getSQLClient),
	}
}
