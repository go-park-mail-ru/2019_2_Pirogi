package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type LikeRepository interface {
	Insert(like model.Like) *model.Error
	Delete(reviewID model.ID, userID model.ID) *model.Error
	Count(reviewID model.ID) int
}
