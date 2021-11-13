package book

import (
	"context"
	"fmt"

	"github.com/cnson19700/book_service/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type pgRepository struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getClient}
}

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.Book, error) {
	db := r.getClient(ctx)
	book := &model.Book{}

	err := db.Where("id = ?", ID).
		First(book).Error

	if err != nil {
		return nil, errors.Wrap(err, "get book by id")
	}

	return book, nil
}

func (r *pgRepository) GetAll(ctx context.Context) ([]model.Book, error) {
	db := r.getClient(ctx)
	listBook := []model.Book{}

	db.Find(&listBook)

	return listBook, nil
}

func (r *pgRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	db := r.getClient(ctx)
	err := db.Create(book).Error

	return book, errors.Wrap(err, "create book")
}

func (r *pgRepository) Delete(ctx context.Context, id int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", id).Delete(&model.Book{}).Error

	return errors.Wrap(err, "delete book fail")
}

func (r *pgRepository) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	db := r.getClient(ctx)
	err := db.Save(book).Error

	return book, errors.Wrap(err, "update book fail")
}

func (r *pgRepository) GetTitle(ctx context.Context, title string) (*model.Book, error) {
	db := r.getClient(ctx)
	book := &model.Book{}

	err := db.Where("title = ?", title).
		First(book).Error

	if err != nil {
		return nil, errors.Wrap(err, "get book by title")
	}

	return book, nil
}

func (r *pgRepository) SearchBook(ctx context.Context,
	paginator *model.Paginator,
	searchText string,
	filter *model.BookFilter,
	orders []string) (*model.BookResult, error) {
	db := r.getClient(ctx)
	query := db.Model(&model.Book{})

	//Order
	for _, order := range orders {
		query.Order(order)
	}

	if filter.CateID != 0 {
		//filterCate = " JOIN book_categories ON book_categories.book_id = books.id AND book_categories.category_id = " + strconv.FormatInt(filter.CateID, 10)
		query.Joins("JOIN book_categories ON book_categories.book_id = books.id AND book_categories.category_id = ?", filter.CateID)
	}

	if filter.AuthorID != 0 {
		//filterAuthor = "AND books.author_id = " + strconv.FormatInt(filter.AuthorID, 10)
		query.Where("books.author_id = ?", filter.AuthorID)
	}

	if filter.MinRating != -1 {
		//filterRate = "AND rating_average > " + strconv.Itoa(filter.MinRating)
		query.Where("books.rating_average > ?", filter.MinRating)
	}

	if searchText != "" {
		//filterTitle = "AND title LIKE " + "'%" + searchText + "%'"
		query.Where("title LIKE ?", "%"+searchText+"%")
	}

	//Paging
	var res model.BookResult

	if paginator.Limit >= 0 {
		if paginator.Page <= 0 {
			paginator.Page = 1
		}

		if paginator.Limit == 0 {
			paginator.Limit = model.PageSize
		}
		res.Page = paginator.Page
		res.Limit = paginator.Limit
		query.Count(&res.Total).Scopes(paginator.Paginate())
	}

	fmt.Print(res)

	err := query.Find(&res.Data).Error

	return &res, err
}

func (r *pgRepository) CreateBookRating(ctx context.Context, bookRating *model.BookRating) (*model.BookRating, error) {
	db := r.getClient(ctx)
	query := db.Model(&model.BookRating{})

	err := query.Create(bookRating).Error

	return bookRating, err
}

func (r *pgRepository) GetAverageBookRating(ctx context.Context, BookID int64) (*model.BookRating, error) {
	db := r.getClient(ctx)
	query := db.Model(&model.BookRating{})

	if BookID != 0 {
		query.Where("book_id = ?", BookID).Group("book_id").Having("count(rating)>1")
	}

	var res model.BookRating
	err := query.Find(&res).Error
	return &res, err
}
