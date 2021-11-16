package book

import (
	"context"
	"fmt"

	"github.com/cnson19700/book_service/model"
	"github.com/cnson19700/book_service/util/myerror"
)

type SearchBookRequest struct {
	Filter    *model.BookFilter
	Paginator *model.Paginator
	OrderBy   string `json:"order_by,omitempty" query:"order_by"`
	OrderType string `json:"order_type,omitempty" query:"order_type"`
}

func (u *Usecase) SearchBook(ctx context.Context, req SearchBookRequest) (*model.BookResult, error) {

	orders := make([]string, 0)
	if req.OrderBy != "" {
		orders = []string{fmt.Sprintf("%s %s", req.OrderBy, req.OrderType)}
	}

	pagnitor := &model.Paginator{
		Limit: 20,
		Page:  1,
	}

	if req.Paginator != nil {
		pagnitor = req.Paginator
	}

	bookList, err := u.bookRepo.SearchBook(ctx, pagnitor, req.Filter, orders)
	if err != nil {
		return nil, myerror.ErrGetBook(err)
	}

	return bookList, nil
}
