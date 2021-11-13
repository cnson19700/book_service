package book

import (
	"context"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/cnson19700/book_service/model"
	checkform "github.com/cnson19700/book_service/package/checkForm"
	imgvalid "github.com/cnson19700/book_service/package/fileValid"
	"github.com/cnson19700/book_service/util/myerror"
)

type UpdateBookRequest struct {
	Form   *multipart.Form
	BookID int64 `json:"book_id"`
}

func (u *Usecase) Update(ctx context.Context, req UpdateBookRequest) (*model.Book, error) {
	book, err := u.bookRepo.GetById(ctx, req.BookID)
	if err != nil {
		return nil, myerror.ErrGetBook(err)
	}

	if len(req.Form.Value["title"][0]) != 0 {
		formTitle := req.Form.Value["title"][0]
		isTitle, title := checkform.CheckFormatValue("title", formTitle)
		if !isTitle {
			return nil, myerror.ErrTitleFormat(nil)
		}

		book.Title = title
	}

	if len(req.Form.Value["language"][0]) != 0 {
		formLanguage := req.Form.Value["language"][0]
		isLanguage, language := checkform.CheckFormatValue("language", formLanguage)
		if !isLanguage {
			return nil, myerror.ErrTitleFormat(nil)
		}

		book.Language = language
	}

	if len(req.Form.Value["description"][0]) != 0 {
		formDescription := req.Form.Value["description"][0]
		isDescription, description := checkform.CheckFormatValue("description", formDescription)
		if !isDescription {
			return nil, myerror.ErrTitleFormat(nil)
		}

		book.Description = description
	}

	if len(req.Form.File["image"]) != 0 {
		file := req.Form.File["image"][0]
		var initImage = "blank.png"
		var pathFile = "/images/bookImages"

		fileType, err := imgvalid.CheckImage(file)
		if err != nil {
			return nil, err
		}

		initImage = strconv.FormatInt(req.BookID, 10) + "." + strings.Split(fileType, "/")[1]

		err = imgvalid.CopyFile(file, initImage, "."+pathFile)
		if err != nil {
			log.Fatal(err)
		}

		book.Image = pathFile + initImage
	}

	if req.Form.Value["rating_average"][0] != "" {
		ratingAvg, err := strconv.ParseFloat(req.Form.Value["rating_average"][0], 64)

		if err != nil {
			return nil, myerror.ErrRatingAvgFormat(nil)
		}

		book.RatingAverage = ratingAvg
	}

	if req.Form.Value["release_date"][0] != "" {
		releaseDate, err := time.Parse(time.RFC3339, req.Form.Value["release_date"][0])

		if err != nil {
			return nil, myerror.ErrReleaseDateFormat(nil)
		}

		book.ReleaseDate = releaseDate
	}
	res, err := u.bookRepo.Update(ctx, book)

	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}
