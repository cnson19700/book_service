package book

import (
	"github.com/cnson19700/book_service/usecase"
	"github.com/cnson19700/book_service/usecase/book"
	"github.com/labstack/echo/v4"
)

type Route struct {
	bookUseCase book.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		bookUseCase: useCase.Book,
	}
	group.POST("", r.Insert)
	group.DELETE("/:id", r.Delete)
	group.PUT("/:id", r.Update)
	group.GET("", r.SearchBook)
	group.GET("/find", r.GetBook)
	group.POST("/rating", r.Rating)
	group.GET("/rating", r.GetAverageBookRating)
}
