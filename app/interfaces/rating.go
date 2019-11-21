package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
)

type ratingRepository struct {
	conn database.Database
}

func NewRatingRepository(conn database.Database) *ratingRepository {
	return &ratingRepository{conn: conn}
}
