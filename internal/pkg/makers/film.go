package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/security"
	"html"
)

func MakeFilm(id models.ID, in *models.NewFilm) models.Film {
	return models.Film{
		ID:          id,
		Title:       html.EscapeString(in.Title),
		Year:        in.Year,
		Genres:      security.XSSFilterGenres(in.Genres),
		Mark:        models.Mark(0),
		Description: html.EscapeString(in.Description),
		Countries:   security.XSSFilterStrings(in.Countries),
		PersonsID:   in.PersonsID,
		Images:      []models.Image{"default.png"},
		ReviewsNum:  0,
		Trailer:     in.Trailer,
	}
}

func MakeFilmTrunc(in models.Film, persons []models.PersonTrunc) models.FilmTrunc {
	return models.FilmTrunc{
		ID:          in.ID,
		Title:       in.Title,
		Year:        in.Year,
		Genres:      in.Genres,
		Mark:        in.Mark,
		Description: in.Description,
		Persons:     persons,
		Image:       in.Images[0],
	}
}

func MakePersonTrunc(in models.Person) models.PersonTrunc {
	return models.PersonTrunc{
		ID:    in.ID,
		Name:  in.Name,
		Image: in.Images[0],
	}
}

func MakeFilmFull(in models.Film, persons []models.Person) models.FilmFull {
	var personsTrunc []models.PersonTrunc
	for _, person := range persons {
		personsTrunc = append(personsTrunc, MakeTruncPerson(person))
	}
	return models.FilmFull{
		ID:          in.ID,
		Title:       in.Title,
		Year:        in.Year,
		Genres:      in.Genres,
		Mark:        in.Mark,
		Description: in.Description,
		Countries:   in.Countries,
		Persons:     personsTrunc,
		Images:      in.Images,
		ReviewsNum:  in.ReviewsNum,
		Trailer:     in.Trailer,
	}
}

func MakeFilmsTrunc(in []models.Film, persons [][]models.PersonTrunc) (out []models.FilmTrunc) {
	for i, film := range in {
		out = append(out, MakeFilmTrunc(film, persons[i]))
	}
	return
}

func MakePersonsTrunc(in []models.Person) (out []models.PersonTrunc) {
	for _, person := range in {
		out = append(out, MakePersonTrunc(person))
	}
	return
}
