package book

import (
	"errors"
	"strconv"

	"github.com/cnson19700/book_service/usecase/book"
	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (r *Route) Update(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)
	bookID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	form, err := c.MultipartForm()
	if err != nil {
		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	user, err := r.bookUseCase.Update(ctx, book.UpdateBookRequest{Form: form, BookID: bookID})
	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, user)
}
