package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type ReviewRepository interface {
	Insert(newReview model.ReviewNew) (model.ID, error)
	Update(id model.ID, review model.Review) error
	Delete(id model.ID) bool
	Get(id model.ID) model.Review
	GetMany(id []model.ID) []model.Review
}
