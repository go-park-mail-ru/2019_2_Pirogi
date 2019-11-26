package modelWorker

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

func MakeTrailersList(films []model.Film) (trailers []model.TrailerWithTitle) {
	for _, film := range films {
		trailer := model.TrailerWithTitle{
			Title:   film.Title,
			Trailer: film.Trailer,
		}
		trailers = append(trailers, trailer)
	}
	return trailers
}
