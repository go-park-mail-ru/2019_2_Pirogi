package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type ListRepository interface {
	Insert(list model.List) *model.Error
	Update(list model.List) *model.Error
	Get(id model.ID) (model.List, *model.Error)
	GetByUserID(userID model.ID) (model.List, *model.Error)
}
