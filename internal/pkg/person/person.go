package person

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

func CreatePerson(id models.ID, in models.Person) (models.Person, *models.Error) {
	return models.Person{
		ID:         id,
		Roles:      in.Roles,
		Name:       in.Name,
		Birthday:   in.Birthday,
		Birthplace: in.Birthplace,
		Genres:     in.Genres,
		Films:      in.Films,
		Likes:      0,
		Images:     in.Images,
		//Awards:     in.Awards,
	}, nil
}
