package book

import (
	"context"

	"github.com/cnson19700/book_service/model"
)

type Repository interface {
	GetById(ctx context.Context, ID int64) (*model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	Delete(ctx context.Context, ID int64) error
	Insert(ctx context.Context, user *model.Book) (*model.Book, error)
	Update(ctx context.Context, user *model.Book) (*model.Book, error)
	SearchBook(ctx context.Context,
		page *model.Paginator,

		filter *model.BookFilter,
		orders []string) (*model.BookResult, error)
	CreateBookRating(ctx context.Context, bookRating *model.BookRating) (*model.BookRating, error)
	GetAverageBookRating(ctx context.Context, BookID int64) (*model.BookRating, error)
}
