package repository

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
)

type UserRepository interface {
	Insert(newUser model.UserNew) *model.Error
	Update(user model.User) *model.Error
	Delete(id model.ID) *model.Error
	Get(id model.ID) (model.User, *model.Error)
	GetMany(ids []model.ID) []model.User
	GetByEmail(email string) (model.User, *model.Error)
}
