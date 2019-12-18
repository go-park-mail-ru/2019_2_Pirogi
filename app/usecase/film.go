package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
)

type FilmUsecase interface {
	Create(body []byte) *model.Error
	List(ids []model.ID) []model.Film
	GetFilmFullByte(id model.ID) ([]byte, *model.Error)
}

type filmUsecase struct {
	filmRepo         repository.FilmRepository
	personRepo       repository.PersonRepository
	subscriptionRepo repository.SubscriptionRepository
}

func NewFilmUsecase(filmRepo repository.FilmRepository, personRepository repository.PersonRepository,
	subscriptionRepository repository.SubscriptionRepository) *filmUsecase {
	return &filmUsecase{
		filmRepo:         filmRepo,
		personRepo:       personRepository,
		subscriptionRepo: subscriptionRepository,
	}
}

func (u *filmUsecase) Create(body []byte) *model.Error {
	filmNew := model.FilmNew{}
	err := filmNew.UnmarshalJSON(body)
	if err != nil {
		return model.NewError(400, "Невалидные данные ", err.Error())
	}
	e := u.filmRepo.Insert(filmNew)
	if e != nil {
		return e
	}
	film, ok := u.filmRepo.GetByTitle(filmNew.Title)
	if !ok {
		return model.NewError(500, "Ошибка при сохранении нового фильма", filmNew.Title)
	}
	id := film.ID
	persons := u.personRepo.GetMany(filmNew.PersonsID)

	for idx, person := range persons {
		if !person.HasFilmID(id) {
			persons[idx].FilmsID = append(person.FilmsID, id)
		}
		for _, filmGenre := range filmNew.Genres {
			if !person.HasGenre(filmGenre) {
				persons[idx].Genres = append(person.Genres, filmGenre)
			}
		}
		go u.personRepo.Update(persons[idx])
		event := model.NewSubscriptionEvent(person.ID, film.ID, "Новый фильм \""+film.Title+"\" с "+person.Name)
		go u.subscriptionRepo.SendEventToSubscribers(event)
	}
	return e
}

func (u *filmUsecase) List(ids []model.ID) []model.Film {
	return u.filmRepo.GetMany(ids)
}

func (u *filmUsecase) GetFilmFullByte(id model.ID) ([]byte, *model.Error) {
	film, err := u.filmRepo.Get(id)
	if err != nil {
		return nil, err
	}
	persons := u.personRepo.GetMany(film.PersonsID)
	filmFull := film.Full(persons)
	related, err := u.filmRepo.GetRelated(filmFull)
	if err != nil {
		return nil, err
	}
	filmFull.Related = related
	body, e := filmFull.MarshalJSON()
	if e != nil {
		return nil, model.NewError(500, e.Error())
	}
	return body, nil
}
