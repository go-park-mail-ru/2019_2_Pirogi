package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
)

func MakeTrailersList(films []domains.Film) (trailers []domains.TrailerWithTitle) {
	for _, film := range films {
		trailer := domains.TrailerWithTitle{
			Title:   film.Title,
			Trailer: film.Trailer,
		}
		trailers = append(trailers, trailer)
	}
	return
}
