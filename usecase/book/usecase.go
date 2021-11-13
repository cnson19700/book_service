package book

import (
	"github.com/cnson19700/book_service/repository"
	"github.com/cnson19700/book_service/repository/book"
)

type Usecase struct {
	bookRepo book.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		bookRepo: repo.Book,
	}
}
