package user

import (
	"context"

	"github.com/cnson19700/pkg/middleware"
	"github.com/cnson19700/user_service/model"
	"github.com/cnson19700/user_service/util/myerror"
)

func (u *Usecase) GetMe(ctx context.Context) (*model.User, error) {
	payload := middleware.GetClaim(ctx)
	id := payload.UserID

	user, err := u.userRepo.GetById(ctx, id)
	if err != nil {
		return &model.User{}, myerror.ErrGetUser(err)
	}

	return user, nil
}
