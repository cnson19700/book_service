package book

import (
	"context"

	"github.com/cnson19700/book_service/model"
	"github.com/cnson19700/book_service/util/myerror"
	"github.com/cnson19700/pkg/middleware"
)

type RatingRequest struct {
	BookID int64 `json:"book_id"`
	Rating int   `json:"rating"`
}

func (u *Usecase) CreateRatingBook(ctx context.Context, req RatingRequest) (*model.BookRating, error) {
	claim := middleware.GetClaim(ctx)
	rating := &model.BookRating{
		UserID: claim.UserID,
		BookID: req.BookID,
		Rating: req.Rating,
	}

	res, err := u.bookRepo.CreateBookRating(ctx, rating)
	if err != nil {
		return &model.BookRating{}, myerror.ErrRatingAvgFormat(err)
	}

	return res, nil
}
