package film

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func CreateFilm(id models.ID, newFilm *models.NewFilm) (models.Film, *models.Error) {
	return models.Film{
		FilmTrunc: models.FilmTrunc{
			ID:     id,
			Title:  newFilm.Title,
			Year:   newFilm.Year,
			Genres: newFilm.Genres,
			Poster: newFilm.Poster,
			Rating: 0,
		},
		Description: newFilm.Description,
		Actors:      newFilm.Actors,
		Directors:   newFilm.Directors,
		Images:      newFilm.Images,
		ReviewsNum:  0,
	}, nil
}
