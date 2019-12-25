package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
)

type ratingRepository struct {
	conn database.Database
}

func (r *ratingRepository) Upsert(rating model.Rating) *model.Error {
	return r.conn.Upsert(rating)
}

func NewRatingRepository(conn database.Database) *ratingRepository {
	return &ratingRepository{conn: conn}
}
