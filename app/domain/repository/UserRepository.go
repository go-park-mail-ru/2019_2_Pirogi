package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type UserRepository interface {
	Insert(newUser model.UserNew) (model.ID, error)
	Update(id model.ID, user model.User) error
	Delete(id model.ID) bool
	Get(id model.ID) model.User
	GetMany(target model.Target, id model.ID) []model.User
	GetByCookie(cookie model.Cookie) (model.User, bool)
	GetByEmail(email string) (model.User, bool)
	MakeTrunc(user model.User) model.UserTrunc
}
