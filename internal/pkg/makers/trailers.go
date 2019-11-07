package makers

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

func MakeTrailersList(films []models.Film) (trailers []string) {
	for _, film := range films {
		trailers = append(trailers, film.Trailer)
	}
	return
}
