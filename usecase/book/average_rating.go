package book

import (
	"context"

	"github.com/cnson19700/book_service/model"
	bookError "github.com/cnson19700/book_service/util/myerror"
)

type GetBookRatingRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) GetAverageBookRating(ctx context.Context, req GetBookRatingRequest) (*model.BookRating, error) {
	res, err := u.bookRepo.GetAverageBookRating(ctx, req.ID)
	if err != nil {
		return &model.BookRating{}, bookError.ErrGetRatingAvg(err)
	}

	return res, nil
}
