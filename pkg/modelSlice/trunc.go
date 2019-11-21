package modelSlice

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

func TruncFilms(films []model.Film) (filmsTrunc []model.FilmTrunc) {
	for _, film := range films {
		filmsTrunc = append(filmsTrunc, film.Trunc())
	}
	return
}


func TruncPersons(persons []model.Person) (personsTrunc []model.PersonTrunc) {
	for _, person := range persons {
		personsTrunc = append(personsTrunc, person.Trunc())
	}
	return
}