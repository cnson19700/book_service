package user

import (
	"context"

	"github.com/cnson19700/auth_service/package/auth"
	"github.com/cnson19700/pkg/middleware"
	checkform "github.com/cnson19700/auth_service/package/checkForm"
	"github.com/cnson19700/user_service/util/myerror"
)

type UpdatePasswordRequest struct {
	Password  string `json:"password"`
	NewPass   string `json:"new_password"`
	ReNewPass string `json:"re_new_password"`
}

func (u *Usecase) UpdatePassword(ctx context.Context, req UpdatePasswordRequest) error {
	payload := middleware.GetClaim(ctx)
	id := payload.UserID

	if req.NewPass != req.ReNewPass {
		return myerror.ErrMatchPassword(nil)
	}

	//password format error
	isPass, newPass := checkform.CheckFormatValue("password", req.NewPass)
	if !isPass {
		return myerror.ErrPasswordFormat(nil)
	}

	isPass, oldPass := checkform.CheckFormatValue("password", req.Password)
	if !isPass {
		return myerror.ErrPasswordFormat(nil)
	}

	user, err := u.userRepo.GetById(ctx, id)
	if err != nil {
		return myerror.ErrGetUser(err)
	}

	isPassTrue := auth.VerifyPassword(oldPass, user.Password)
	if !isPassTrue {
		return myerror.ErrInvalidPassword(err)
	}

	newPassHash, err := auth.HashPassword(newPass)

	if err != nil {
		return myerror.ErrHashPassword(err)
	}

	user.Password = newPassHash

	_, err = u.userRepo.Update(ctx, user)
	if err != nil {
		return myerror.ErrUpdatePassword(err)
	}

	return nil
}
