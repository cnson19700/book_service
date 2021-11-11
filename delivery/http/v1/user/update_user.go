package user

import (
	"errors"

	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/cnson19700/user_service/usecase/user"
	"github.com/labstack/echo/v4"
)

func (r *Route) UpdateUser(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)
	form, err := c.MultipartForm()
	if err != nil {
		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	user, err := r.userUseCase.UpdateUser(ctx, user.UpdateRequest{Form: form})
	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, user)
}
