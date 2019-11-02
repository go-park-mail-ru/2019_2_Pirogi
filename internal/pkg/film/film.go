package film

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	person2 "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/person"
	"github.com/labstack/gommon/log"
)

func CreateFilm(id models.ID, in *models.NewFilm) models.Film {
	return models.Film{
		ID:          id,
		Title:       in.Title,
		Year:        in.Year,
		Genres:      in.Genres,
		Poster:      "",
		Mark:        models.Mark(0),
		Description: in.Description,
		Countries:   in.Countries,
		PersonsID:   in.PersonsID,
		Images:      nil,
		ReviewsNum:  0,
	}
}

func MakeTruncFilm(in models.Film) models.FilmTrunc {
	return models.FilmTrunc{
		ID:     in.ID,
		Title:  in.Title,
		Year:   in.Year,
		Genres: in.Genres,
		Poster: in.Poster,
		Mark:   in.Mark,
	}
}

func MakerFullFilm(conn database.Database, in models.Film) models.FilmFull {
	var personsTrunc []models.PersonTrunc
	for _, personID := range in.PersonsID {
		tmp, err := conn.Get(personID, "person")
		if err != nil {
			log.Warn(err)
			continue
		}
		person, ok := tmp.(models.Person)
		if !ok {
			log.Warn("MakeFullFilm: can not cast")
			continue
		}
		personsTrunc = append(personsTrunc, person2.MakeTruncPerson(person))
	}
	return models.FilmFull{
		ID:          in.ID,
		Title:       in.Title,
		Year:        in.Year,
		Genres:      in.Genres,
		Poster:      in.Poster,
		Mark:        in.Mark,
		Description: in.Description,
		Countries:   in.Countries,
		Persons:     personsTrunc,
		Images:      in.Images,
		ReviewsNum:  in.ReviewsNum,
	}
}
