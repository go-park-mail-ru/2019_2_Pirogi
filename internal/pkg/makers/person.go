package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/security"
	"html"
)

func MakePerson(id domains.ID, in domains.NewPerson) domains.Person {
	return domains.Person{
		ID:         id,
		Name:       html.EscapeString(in.Name),
		Roles:      security.XSSFilterRoles(in.Roles),
		Birthday:   html.EscapeString(in.Birthday),
		Birthplace: html.EscapeString(in.Birthplace),
		Genres:     []domains.Genre{},
		FilmsID:    []domains.ID{},
		Likes:      0,
		Images:     []domains.Image{"default.png"},
	}
}

func MakeTruncPerson(in domains.Person) domains.PersonTrunc {
	return domains.PersonTrunc{
		ID:    in.ID,
		Name:  in.Name,
		Image: in.Images[0],
	}
}

func MakeFullPerson(in domains.Person, films []domains.Film) domains.PersonFull {
	var filmsTrunc []domains.FilmTrunc
	for _, film := range films {
		filmsTrunc = append(filmsTrunc, MakeFilmTrunc(film))
	}
	return domains.PersonFull{
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
