package usecase

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelSlice"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"strconv"
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
}

func NewReviewUsecase(reviewRepo repository.ReviewRepository, cookieRepo repository.CookieRepository,
	userRepo repository.UserRepository) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo: reviewRepo,
		cookieRepo: cookieRepo,
		userRepo:   userRepo,
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
	body := modelSlice.MarshalReviews(reviews)
	jsonBody := json.MakeJSONArray(body)
	return jsonBody, nil
}

func (u *reviewUsecase) GetFilmReviewsFullJSONBlob(filmID model.ID, limit, offset int) ([]byte, *model.Error) {
	reviews, err := u.reviewRepo.GetMany(configs.Default.FilmTargetName, filmID, limit, offset)
	zap.S().Debug(reviews)
	if err != nil {
		return nil, err
	}
	var reviewsFull[] model.ReviewFull
	for _, review := range reviews {
		user, err := u.userRepo.Get(review.AuthorID)
		if err != nil {
			continue
		}
		reviewsFull = append(reviewsFull, review.Full(user.Trunc()))
	}
	zap.S().Debug(reviewsFull)
	body := modelSlice.MarshalReviewsFull(reviewsFull)
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
