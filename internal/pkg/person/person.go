package person

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
		ImagesID:   []models.ID{},
	}, nil
}
