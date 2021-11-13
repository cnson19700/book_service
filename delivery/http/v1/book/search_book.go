package book

import (
	"errors"
	"strconv"

	"github.com/cnson19700/book_service/model"
	"github.com/cnson19700/book_service/usecase/book"
	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (r *Route) SearchBook(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)
	if err := c.Bind(&model.Paginator{}); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	searchText := c.QueryParam("search")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	authorID, _ := strconv.ParseInt(c.QueryParam("author_id"), 10, 64)
	cateID, _ := strconv.ParseInt(c.QueryParam("cate_id"), 10, 64)
	minRating, _ := strconv.Atoi(c.QueryParam("min_rating"))

	req := book.SearchBookRequest{
		Paginator: &model.Paginator{
			Page:  page,
			Limit: limit,
		},
		Filter: &model.BookFilter{
			AuthorID:  authorID,
			CateID:    cateID,
			MinRating: minRating,
		},
	}

	//log.Fatal(req)

	res, err := r.bookUseCase.SearchBook(ctx, searchText, req)
	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, res)
}
