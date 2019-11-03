package makers

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

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
		Images:     []models.Image{"default.png"},
	}, nil
}

func MakeTruncPerson(in models.Person) models.PersonTrunc {
	return models.PersonTrunc{
		ID:   in.ID,
		Name: in.Name,
	}
}

func MakeFullPerson(in models.Person, films []models.Film) models.PersonFull {
	var filmsTrunc []models.FilmTrunc
	for _, film := range films {
		filmsTrunc = append(filmsTrunc, MakeTruncFilm(film))
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
