package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/labstack/echo"
	"net/http"
)

type RatingUsecase interface {
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
	CreateOrUpdateRating(body []byte, user model.User) *model.Error
}

type ratingUsecase struct {
	ratingRepo repository.RatingRepository
	cookieRepo repository.CookieRepository
	filmRepo   repository.FilmRepository
}

func NewRatingUsecase(ratingRepo repository.RatingRepository, cookieRepo repository.CookieRepository,
	filmRepo repository.FilmRepository) *ratingUsecase {
	return &ratingUsecase{
		ratingRepo: ratingRepo,
		cookieRepo: cookieRepo,
		filmRepo:   filmRepo,
	}
}

func (u *ratingUsecase) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	return u.cookieRepo.GetUserByContext(ctx)
}

func (u *ratingUsecase) CreateOrUpdateRating(body []byte, user model.User) *model.Error {
	var ratingNew model.RatingNew
	err := ratingNew.UnmarshalJSON(body)
	if err != nil {
		return model.NewError(http.StatusNotFound, "Невалидные данные ", err.Error())
	}

	film, e := u.filmRepo.Get(ratingNew.FilmID)
	if e != nil {
		return e
	}

	foundRating, e := u.ratingRepo.FindRatingByUserIDAndFilmID(user.ID, ratingNew.FilmID)
	if e != nil {
		rating := ratingNew.ToRating(user.ID)
		e = u.ratingRepo.Insert(rating)
		if e != nil {
			return e
		}
		film.RatingSum += int(rating.Stars)
		film.VotersNum++
	} else {
		ratingUpdate := foundRating.ToRatingUpdate()
		ratingUpdate.Stars = ratingNew.Stars
		e = u.ratingRepo.Update(ratingUpdate)
		if e != nil {
			return e
		}
		film.RatingSum = film.RatingSum - int(foundRating.Stars) + int(ratingUpdate.Stars)
	}

	film.CountAndSetMark()
	u.filmRepo.Update(film)
	return nil
}
