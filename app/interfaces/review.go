package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
)

type reviewRepository struct {
	conn database.Database
}

func (r *reviewRepository) Insert(newReview model.ReviewNew) *model.Error {
	return r.conn.Upsert(newReview)
}

func (r *reviewRepository) Update(id model.ID, review model.Review) *model.Error {
	panic("implement me")
}

func (r *reviewRepository) Delete(id model.ID) *model.Error {
	panic("implement me")
}

func (r reviewRepository) GetMany(target string, id model.ID, limit, offset int) (reviews []model.Review, err *model.Error) {
	switch target {
	case configs.Default.FilmTargetName:
		return r.conn.GetReviewsOfFilmSortedByDate(id, limit, offset)
	case configs.Default.UserTargetName:
		return r.conn.GetReviewsOfAuthorSortedByDate(id, limit, offset)
	default:
		return nil, model.NewError(400, "bad target")
	}
}

func NewReviewRepository(conn database.Database) *reviewRepository {
	return &reviewRepository{
		conn: conn,
	}
}
