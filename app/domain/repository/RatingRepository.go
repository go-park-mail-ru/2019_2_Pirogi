package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type RatingRepository interface {
	Upsert(rating model.Rating) *model.Error
	//GetRatingByUserAndFilm(userId model.ID, filmId model.ID) (model.Rating, *model.Error)
}
