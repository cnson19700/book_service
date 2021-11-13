package book

import (
	"context"

	"github.com/cnson19700/book_service/util/myerror"
)

type DeleteBookRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) Delete(ctx context.Context, req DeleteBookRequest) error {
	book, err := u.bookRepo.GetById(ctx, req.ID)
	if err != nil {
		return myerror.ErrGetBook(err)
	}

	err = u.bookRepo.Delete(ctx, book.ID)

	return err
}
