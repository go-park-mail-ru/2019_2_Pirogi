package person

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/gommon/log"
)

func CreatePerson(id models.ID, in models.NewPerson) (models.Person, *models.Error) {
	return models.Person{
		ID:         id,
		Name:       in.Name,
		Roles:      in.Roles,
		Birthday:   in.Birthday,
		Birthplace: in.Birthplace,
		Genres:     []models.Genre{},
		FilmsID:    []models.ID{},
		Likes:      0,
		Images:     []models.Image{},
	}, nil
}

func MakeTruncPerson(in models.Person) models.PersonTrunc {
	return models.PersonTrunc{
		ID:   in.ID,
		Name: in.Name,
		Mark: in.Mark,
	}
}

func MakeFullPerson(conn database.Database, in models.Person) models.PersonFull {
	var filmsTrunc []models.FilmTrunc
	for _, filmID := range in.FilmsID {
		tmp, err := conn.Get(filmID, "film")
		if err != nil {
			log.Warn("Make full person error: ", err)
			continue
		}
		filmTrunc, ok := tmp.(models.Film)
		if !ok {
			log.Warn("Make full person error: can not cast into type FilmTrunc")
			continue
		}
		filmsTrunc = append(filmsTrunc, film.MakeTruncFilm(filmTrunc))
	}
	return models.PersonFull{
		ID:         in.ID,
		Name:       in.Name,
		Mark:       in.Mark,
		Roles:      in.Roles,
		Birthday:   in.Birthday,
		Birthplace: in.Birthplace,
		Genres:     in.Genres,
		Films:      filmsTrunc,
		Likes:      in.Likes,
		Images:     in.Images,
	}
}
