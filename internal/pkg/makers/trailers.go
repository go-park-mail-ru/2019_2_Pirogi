package makers

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

func MakeTrailersList(films []models.Film) (trailers []models.TrailerWithTitle) {
	for _, film := range films {
		trailer := models.TrailerWithTitle{
			Title:   film.Title,
			Trailer: film.Trailer,
		}
		trailers = append(trailers, trailer)
	}
	return
}
