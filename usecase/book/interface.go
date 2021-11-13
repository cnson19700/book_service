package book

import (
	"context"

	"github.com/cnson19700/book_service/model"
)

type IUsecase interface {
	Insert(ctx context.Context, req InsertRequest) (*model.Book, error)
	Delete(ctx context.Context, req DeleteBookRequest) error
	Update(ctx context.Context, req UpdateBookRequest) (*model.Book, error)
	SearchBook(ctx context.Context, searchText string, req SearchBookRequest) (*model.BookResult, error)
	GetBook(ctx context.Context, req GetBookRequest) (*model.Book, error)
	CreateRatingBook(ctx context.Context, req RatingRequest) (*model.BookRating, error)
	GetAverageBookRating(ctx context.Context, req GetBookRatingRequest) (*model.BookRating, error)
}
