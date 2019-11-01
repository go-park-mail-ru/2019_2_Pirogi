package person

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

func CreatePerson(id models.ID, in models.NewPerson) (models.Person, *models.Error) {
	return models.Person{
		PersonTrunc: models.PersonTrunc{
			ID:   id,
			Name: in.Name,
		},
		Roles:      in.Roles,
		Birthday:   in.Birthday,
		Birthplace: in.Birthplace,
		Genres:     in.Genres,
		Films:      in.Films,
		Likes:      0,
		Images:     in.Images,
	}, nil
}
