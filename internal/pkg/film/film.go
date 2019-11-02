package film

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func CreateFilm(id models.ID, newFilm *models.NewFilm) (models.Film, *models.Error) {
	return models.Film{
		ID:          id,
		Title:       newFilm.Title,
		Year:        newFilm.Year,
		Genres:      newFilm.Genres,
		Mark:        0,
		Description: newFilm.Description,
		PersonsID:   newFilm.PersonsID,
		ImagesID:    []models.ID{},
		ReviewsNum:  0,
	}, nil
}
