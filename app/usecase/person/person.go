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
	return u.personRepo.Insert(personNew)
}

func (u *personUsecase) List(ids []model.ID) []model.Person {
	return u.personRepo.GetMany(ids)
}

func (u *personUsecase) GetPersonFullByte(id model.ID) ([]byte, *model.Error) {
	person, err := u.personRepo.Get(id)
	if err != nil {
		return nil, err
	}
	films := u.filmRepo.GetMany(person.FilmsID)
	personFull := person.Full(films)
	body, e := personFull.MarshalJSON()
	if e != nil {
		return nil, model.NewError(500, e.Error())
	}
	return body, nil
}
