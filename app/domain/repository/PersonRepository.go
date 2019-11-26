package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type PersonRepository interface {
	Insert(newPerson model.PersonNew) *model.Error
	Update(person model.Person) *model.Error
	Delete(id model.ID) bool
	Get(id model.ID) (model.Person, *model.Error)
	GetMany(id []model.ID) []model.Person
	MakeTrunc(person model.Person) model.PersonTrunc
	MakeFull(person model.Person) model.PersonFull
	GetByPipeline(pipeline interface{}) ([]model.Person, *model.Error)
}
