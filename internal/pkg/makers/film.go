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

func MakeFilmTrunc(in models.Film) models.FilmTrunc {
	return models.FilmTrunc{
		ID:     in.ID,
		Title:  in.Title,
		Year:   in.Year,
		Genres: in.Genres,
		Mark:   in.Mark,
		Image:  in.Images[0],
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

func MakeFilmsTrunc(in []models.Film) (out []models.FilmTrunc) {
	for _, film := range in {
		out = append(out, MakeFilmTrunc(film))
	}
	return
}
