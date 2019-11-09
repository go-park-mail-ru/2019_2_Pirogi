package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/security"
	"html"
)

func MakePerson(id models.ID, in models.NewPerson) models.Person {
	return models.Person{
		ID:         id,
		Name:       html.EscapeString(in.Name),
		Roles:      security.XSSFilterRoles(in.Roles),
		Birthday:   html.EscapeString(in.Birthday),
		Birthplace: html.EscapeString(in.Birthplace),
		Genres:     []models.Genre{},
		FilmsID:    []models.ID{},
		Likes:      0,
		Images:     []models.Image{"default.png"},
	}
}

func MakeTruncPerson(in models.Person) models.PersonTrunc {
	var image models.Image
	if len(in.Images) == 0 {
		image = ""
	} else {
		image = in.Images[0]
	}
	return models.PersonTrunc{
		ID:    in.ID,
		Name:  in.Name,
		Image: image,
	}
}

func MakeFullPerson(in models.Person, films []models.Film) models.PersonFull {
	var filmsTrunc []models.FilmTrunc
	for _, film := range films {
		filmsTrunc = append(filmsTrunc, MakeFilmTrunc(film))
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
