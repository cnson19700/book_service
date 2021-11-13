package book

import (
	"errors"
	"strconv"

	"github.com/cnson19700/book_service/usecase/book"
	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	// Bind order by
	if err := c.Bind(&book.DeleteBookRequest{}); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	bookID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	//log.Fatal(bookID)

	err := r.bookUseCase.Delete(ctx, book.DeleteBookRequest{ID: bookID})

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, nil)
}
