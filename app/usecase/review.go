package usecase

import (
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelWorker"
	"github.com/labstack/echo"
)

type ReviewUsecase interface {
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
	GetUserReviewsJSONBlob(user model.User, offset, limit int) ([]byte, *model.Error)
	GetLimitAndOffset(ctx echo.Context) (limit, offset int)
	GetFilmReviewsFullJSONBlob(filmID model.ID, offset, limit int) ([]byte, *model.Error)
	CreateNewReview(body []byte, user model.User) *model.Error
}

type reviewUsecase struct {
	reviewRepo repository.ReviewRepository
	cookieRepo repository.CookieRepository
	userRepo   repository.UserRepository
	filmRepo   repository.FilmRepository
}

func NewReviewUsecase(reviewRepo repository.ReviewRepository, cookieRepo repository.CookieRepository,
	userRepo repository.UserRepository, filmRepo repository.FilmRepository) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo: reviewRepo,
		cookieRepo: cookieRepo,
		userRepo:   userRepo,
		filmRepo:   filmRepo,
	}
}

func (u *reviewUsecase) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	return u.cookieRepo.GetUserByContext(ctx)
}

func (u *reviewUsecase) GetUserReviewsJSONBlob(user model.User, limit, offset int) ([]byte, *model.Error) {
	reviews, err := u.reviewRepo.GetMany(configs.Default.UserTargetName, user.ID, limit, offset)
	if err != nil {
		return nil, err
	}

	// TODO: неэффективно, переписать на один запрос к бд
	for i, review := range reviews {
		film, err := u.filmRepo.Get(review.FilmID)
		if err != nil {
			return nil, model.NewError(500, err.Error)
		}
		reviews[i].FilmTitle = film.Title
	}

	body := modelWorker.MarshalReviews(reviews)
	jsonBody := json.MakeJSONArray(body)
	return jsonBody, nil
}

func (u *reviewUsecase) GetFilmReviewsFullJSONBlob(filmID model.ID, limit, offset int) ([]byte, *model.Error) {
	reviews, err := u.reviewRepo.GetMany(configs.Default.FilmTargetName, filmID, limit, offset)
	if err != nil {
		return nil, err
	}
	var reviewsFull []model.ReviewFull
	for _, review := range reviews {
		user, err := u.userRepo.Get(review.AuthorID)
		if err != nil {
			continue
		}
		reviewsFull = append(reviewsFull, review.Full(user.Trunc()))
	}
	body := modelWorker.MarshalReviewsFull(reviewsFull)
	jsonBody := json.MakeJSONArray(body)
	return jsonBody, nil
}

func (u *reviewUsecase) GetLimitAndOffset(ctx echo.Context) (limit, offset int) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		limit = 10
	}
	offset, err = strconv.Atoi(ctx.Param("offset"))
	if err != nil {
		offset = 0
	}
	return limit, offset
}

func (u *reviewUsecase) CreateNewReview(body []byte, user model.User) *model.Error {
	var newReview model.ReviewNew
	err := newReview.UnmarshalJSON(body)
	if err != nil {
		return model.NewError(400, err.Error())
	}
	_, err = govalidator.ValidateStruct(newReview)
	if err != nil {
		return model.NewError(400, err.Error())
	}
	newReview.AuthorID = user.ID
	e := u.reviewRepo.Insert(newReview)
	if e != nil {
		return e
	}
	return nil
}
