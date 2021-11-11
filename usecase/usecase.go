package usecase

import (
	"github.com/cnson19700/user_service/repository"
	"github.com/cnson19700/user_service/usecase/user"
)

type UseCase struct {
	User user.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{

		User: user.New(repo),
	}
}
