package film

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func CreateNewFilm(id models.ID, newFilm *models.NewFilm) (models.Film, *models.Error) {
	return models.Film{
		ID:       id,
		FilmInfo: newFilm.FilmInfo,
	}, nil
}
