package makers

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

func MakeFilm(id models.ID, in *models.NewFilm) models.Film {
	return models.Film{
		ID:          id,
		Title:       in.Title,
		Year:        in.Year,
		Genres:      in.Genres,
		Mark:        models.Mark(0),
		Description: in.Description,
		Countries:   in.Countries,
		PersonsID:   in.PersonsID,
		Images:      []models.Image{"default.png"},
		ReviewsNum:  0,
	}
}

func MakeTruncFilm(in models.Film) models.FilmTrunc {
	return models.FilmTrunc{
		ID:     in.ID,
		Title:  in.Title,
		Year:   in.Year,
		Genres: in.Genres,
		Mark:   in.Mark,
	}
}

func MakeFullFilm(in models.Film, persons []models.Person) models.FilmFull {
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
	}
}
