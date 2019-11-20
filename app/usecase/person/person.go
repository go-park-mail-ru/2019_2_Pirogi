package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
)

type PersonUsecase interface {
	Create(body []byte) *model.Error
	List(ids []model.ID) []model.Person
	GetPersonFullByte(id model.ID) ([]byte, *model.Error)
}

type personUsecase struct {
	personRepo repository.PersonRepository
	filmRepo   repository.FilmRepository
}

func NewPersonUsecase(personRepo repository.PersonRepository, filmRepo repository.FilmRepository) *personUsecase {
	return &personUsecase{
		personRepo: personRepo,
		filmRepo:   filmRepo,
	}
}

func (u *personUsecase) Create(body []byte) *model.Error {
	personNew := model.PersonNew{}
	err := personNew.UnmarshalJSON(body)
	if err != nil {
		return model.NewError(400, "Person: Create: ", err.Error())
	}
	_, e := u.personRepo.Insert(personNew)
	return e
}

func (u *personUsecase) List(ids []model.ID) []model.Person {
	return u.personRepo.GetMany(ids)
}

func (u *personUsecase) GetPersonFullByte(id model.ID) ([]byte, *model.Error) {
	person := u.personRepo.Get(id)
	films := u.filmRepo.GetMany(person.FilmsID)
	personFull := person.Full(films)
	body, err := personFull.MarshalJSON()
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}
	return body, nil
}
