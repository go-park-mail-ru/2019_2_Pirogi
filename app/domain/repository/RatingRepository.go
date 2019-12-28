package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type RatingRepository interface {
	Insert(rating model.Rating) *model.Error
	Update(rating model.RatingUpdate) *model.Error
	FindRatingByUserIDAndFilmID(userID model.ID, filmID model.ID) (model.Rating, *model.Error)
}
