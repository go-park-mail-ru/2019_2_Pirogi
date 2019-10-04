package film

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func CreateNewFilm(id int, newFilm models.NewFilm) (models.Film, *models.Error) {
	film := models.Film{
		ID:       id,
		FilmInfo: newFilm.FilmInfo,
	}
	return film, nil
}
