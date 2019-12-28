package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
)

type ratingRepository struct {
	conn database.Database
}

func (r *ratingRepository) Insert(rating model.Rating) *model.Error {
	return r.conn.Upsert(rating)
}

func (r *ratingRepository) Update(rating model.RatingUpdate) *model.Error {
	return r.conn.Upsert(rating)
}

func (r *ratingRepository) FindRatingByUserIDAndFilmID(userID model.ID, filmID model.ID) (model.Rating, *model.Error) {
	return r.conn.FindRatingByUserIDAndFilmID(userID, filmID)
}

func NewRatingRepository(conn database.Database) *ratingRepository {
	return &ratingRepository{conn: conn}
}
