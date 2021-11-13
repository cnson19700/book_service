package book

import (
	"context"

	"github.com/cnson19700/book_service/model"
	"github.com/cnson19700/book_service/util/myerror"
)

type GetBookRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) GetBook(ctx context.Context, req GetBookRequest) (*model.Book, error) {
	book, err := u.bookRepo.GetById(ctx, req.ID)
	if err != nil {
		return nil, myerror.ErrGetBook(err)
	}

	return book, err
}
