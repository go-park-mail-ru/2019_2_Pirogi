package usecase

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/labstack/echo"
)

type RatingUsecase interface {
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
	CreateOrUpdateRating(body []byte, user model.User) *model.Error
}

type ratingUsecase struct {
	ratingRepo repository.RatingRepository
	cookieRepo repository.CookieRepository
	userRepo   repository.UserRepository
}

func NewRatingUsecase(ratingRepo repository.RatingRepository, cookieRepo repository.CookieRepository,
	userRepo repository.UserRepository) *ratingUsecase {
	return &ratingUsecase{
		ratingRepo: ratingRepo,
		cookieRepo: cookieRepo,
		userRepo:   userRepo,
	}
}

func (u *ratingUsecase) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	return u.cookieRepo.GetUserByContext(ctx)
}

func (u *ratingUsecase) CreateOrUpdateRating(body []byte, user model.User) *model.Error {
	var ratingNew model.RatingNew
	err := ratingNew.UnmarshalJSON(body)
	if err != nil {
		return model.NewError(400, "Невалидные данные ", err.Error())
	}
	_, err = govalidator.ValidateStruct(ratingNew)
	if err != nil {
		return model.NewError(400, "Невалидные данные ", err.Error())
	}
	rating := ratingNew.ToRating(user.ID)
	e := u.ratingRepo.Upsert(rating)
	if e != nil {
		return e
	}
	return nil
}