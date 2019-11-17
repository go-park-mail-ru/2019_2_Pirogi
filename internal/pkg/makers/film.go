package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/security"
	"html"
)

func MakeFilm(id domains.ID, in *domains.NewFilm) domains.Film {
	return domains.Film{
		ID:          id,
		Title:       html.EscapeString(in.Title),
		Year:        in.Year,
		Genres:      security.XSSFilterGenres(in.Genres),
		Mark:        domains.Mark(0),
		Description: html.EscapeString(in.Description),
		Countries:   security.XSSFilterStrings(in.Countries),
		PersonsID:   in.PersonsID,
		Images:      []domains.Image{"default.png"},
		ReviewsNum:  0,
		Trailer:     in.Trailer,
	}
}

func MakeFilmTrunc(in domains.Film) domains.FilmTrunc {
	return domains.FilmTrunc{
		ID:     in.ID,
		Title:  in.Title,
		Year:   in.Year,
		Genres: in.Genres,
		Mark:   in.Mark,
		Image:  in.Images[0],
	}
}

func MakeFilmFull(in domains.Film, persons []domains.Person) domains.FilmFull {
	var personsTrunc []domains.PersonTrunc
	for _, person := range persons {
		personsTrunc = append(personsTrunc, MakeTruncPerson(person))
	}
	return domains.FilmFull{
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

func MakeFilmsTrunc(in []domains.Film) (out []domains.FilmTrunc) {
	for _, film := range in {
		out = append(out, MakeFilmTrunc(film))
	}
	return
}
