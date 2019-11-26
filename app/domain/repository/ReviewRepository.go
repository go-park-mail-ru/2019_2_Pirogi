package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type ReviewRepository interface {
	Insert(newReview model.ReviewNew) *model.Error
	Update(id model.ID, review model.Review) *model.Error
	Delete(id model.ID) *model.Error
	GetMany(target string, id model.ID, limit, offset int) ([]model.Review, *model.Error)
}
